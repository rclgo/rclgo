package rclgo

import (
	"fmt"
	"unsafe"

	cwrap "github.com/rclgo/rclgo/internal"
)

var executor Executor

//
type RosMessage interface {
	Pointer() unsafe.Pointer
	// RosidlMessageTypeSupport() *cwrap.RosidlMessageTypeSupport
	TypeSupport() MessageTypeSupport
}

//
func GetGlobalExecutor() Executor {
	if executor.nodes == nil {
		var err error
		executor, err = NewExecutor()
		if err != nil {
			fmt.Println("NewExecutor error: " + err.Error())
			panic(err)
		}
	}
	return executor
}

//
func Spin(node *Node) {
	exe := GetGlobalExecutor()
	exe.AddNode(node)
	for exe.context.IsValid() {
		exe.SpinOnce()
	}
}

//
func GetReadySubscriptions(ws WaitSet) map[*cwrap.RclSubscription]struct{} {
	var entitiesReady = make(map[*cwrap.RclSubscription]struct{})
	var subs []*cwrap.RclSubscription = ws.rclWaitSet.GetSubscriptions()
	for _, sub := range subs {
		if sub != nil {
			entitiesReady[sub] = struct{}{}
		}
	}
	return entitiesReady
}

// Shutdown the default context.
func Shutdown() error {
	ctx := GetDefaultContext()
	return ctx.Shutdown()
}
