package rclgo

import (
	"fmt"

	cwrap "github.com/rclgo/rclgo/internal"
)

type callable struct {
	handler func()
	entity  interface{}
	node    *Node
}

type Executor struct {
	context Context
	nodes   map[*Node]struct{}
	guard   GuardCondition
}

//
func NewExecutor() (Executor, error) {
	gc, err := NewGuardCondition()
	if err != nil {
		return Executor{}, err
	}
	return Executor{
		context: GetDefaultContext(),
		nodes:   make(map[*Node]struct{}),
		guard:   gc,
	}, nil
}

//
func (e *Executor) AddNode(node *Node) bool {
	_, ok := e.nodes[node]
	if !ok {
		e.nodes[node] = struct{}{}
		node.SetExecutor(e)
		e.guard.Trigger()
		return true
	}
	return false
}

//
func (e *Executor) RemoveNode(node *Node) {
	_, ok := e.nodes[node]
	if ok {
		delete(e.nodes, node)
		e.guard.Trigger()
	}
}

//
func (e *Executor) waitForReadyCallbacks(timeout int64, nodes map[*Node]struct{}) ([]callable, error) {
	nodesToUse := nodes
	if len(nodes) == 0 {
		nodesToUse = e.nodes
	}

	var subscriptions []*cwrap.RclSubscription
	for node := range nodesToUse {
		for sub := range node.subscriptions {
			subscriptions = append(subscriptions, sub.rclSubscription)
		}
	}

	var guards []*cwrap.RclGuardCondition
	var timers []*cwrap.RclTimer
	var clients []*cwrap.RclClient
	var services []*cwrap.RclService
	var events []*cwrap.RclEvent

	waitset := NewWaitSet()
	err := waitset.Init(
		uint(len(subscriptions)),
		uint(len(guards)),
		uint(len(timers)),
		uint(len(clients)),
		uint(len(services)),
		uint(len(events)),
		e.context.rclContext,
		*newDefaultAllocator().rclAllocator,
	)
	if err != nil {
		panic(err)
	}
	err = waitset.Clear()
	if err != nil {
		panic(err)
	}
	for _, sub := range subscriptions {
		err = waitset.AddSubscription(sub)
		if err != nil {
			panic(err)
		}
	}

	err = waitset.Wait(timeout)
	if err != nil {
		return nil, err
	}

	subsReady := GetReadySubscriptions(waitset)

	err = waitset.Fini()
	if err != nil {
		return nil, err
	}

	var ret []callable
	for node := range nodesToUse {
		for sub := range node.subscriptions {
			if _, ok := subsReady[sub.rclSubscription]; ok {
				arg, _, err := e.takeSubscription(sub)
				if err != nil {
					fmt.Printf("e.takeSubscription error: %v\n", err)
					continue
				}
				e.guard.Trigger()
				ret = append(ret, callable{handler: func() { sub.callback(arg) }, entity: sub, node: node})
			}
		}
	}

	return ret, nil
}

//
func (e *Executor) takeSubscription(sub *Subscription) (RosMessage, MessageInfo, error) {
	var msgInfo MessageInfo
	ret := cwrap.RclTake(sub.rclSubscription, sub.msg.Pointer(), msgInfo.RmwMessageInfo, nil)
	if ret != cwrap.Ok {
		return sub.msg, msgInfo, NewErr("RclTake", ret)
	}
	return sub.msg, msgInfo, nil
}

// SpinOnce Single Threaded
func (e *Executor) SpinOnce() error {
	callables, err := e.waitForReadyCallbacks(-1, e.nodes)
	if err != nil {
		return fmt.Errorf("Executor.waitForReadyCallbacks error:\n %s", err.Error())
	}
	for _, callable := range callables {
		callable.handler()
	}
	return nil
}
