package rclgo

import (
	cwrap "github.com/rclgo/rclgo/internal"
)

//
type Subscription struct {
	rclSubscription *cwrap.RclSubscription
	msg             RosMessage
	callback        func(RosMessage)
}

type SubscriptionOptions struct {
	rclSubscriptionOptions *cwrap.RclSubscriptionOptions
}

func newZeroInitializedSubscription() Subscription {
	zeroSubscription := cwrap.RclGetZeroInitializedSubscription()
	return Subscription{rclSubscription: &zeroSubscription}
}

func NewSubscriptionDefaultOptions() SubscriptionOptions {
	defOpts := cwrap.RclSubscriptionGetDefaultOptions()
	return SubscriptionOptions{&defOpts}
}

func (s *Subscription) Init(
	subscriptionOptions SubscriptionOptions,
	node *Node,
	topicName string,
	typeSupport MessageTypeSupport,
) error {

	ret := cwrap.RclSubscriptionInit(
		s.rclSubscription,
		node.rclNode,
		typeSupport.RosidlMessageTypeSupport,
		topicName,
		subscriptionOptions.rclSubscriptionOptions,
	)

	if ret != cwrap.Ok {
		return NewErr("RclSubscriptionInit", ret)
	}

	return nil
}

func (s *Subscription) Fini(node Node) error {
	ret := cwrap.RclSubscriptionFini(
		s.rclSubscription,
		node.rclNode,
	)

	if ret != cwrap.Ok {
		return NewErr("RclSubscriptionFini", ret)
	}

	return nil
}

//
func (s *Subscription) TakeMessage(msgInfo MessageInfo, data MessageData) error {
	if s.rclSubscription == nil || data.Data == nil {
		return NewErr("nil", cwrap.Error)
	}

	ret := cwrap.RclTake(s.rclSubscription, data.Data, msgInfo.RmwMessageInfo, nil)

	if ret != cwrap.Ok {
		return NewErr("RclTake", ret)
	}

	return nil
}
