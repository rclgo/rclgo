package rclgo

import (
	cwrap "github.com/rclgo/rclgo/internal"
)

type Publisher struct {
	rclPublisher *cwrap.RclPublisher
}

type PublisherOptions struct {
	rclPublisherOptions *cwrap.RclPublisherOptions
}

func newZeroInitializedPublisher() Publisher {
	zeroPublisher := cwrap.RclGetZeroInitializedPublisher()
	return Publisher{&zeroPublisher}
}

func NewPublisherDefaultOptions() PublisherOptions {
	defOpts := cwrap.RclPublisherGetDefaultOptions()
	return PublisherOptions{&defOpts}
}

func (p *Publisher) GetTopicName() string {
	return cwrap.RclPublisherGetTopicName(p.rclPublisher)
}

func (p *Publisher) Init(
	publisherOptions PublisherOptions,
	node *Node,
	topicName string,
	typeSupport MessageTypeSupport,
) error {

	ret := cwrap.RclPublisherInit(
		p.rclPublisher,
		node.rclNode,
		typeSupport.RosidlMessageTypeSupport,
		topicName,
		publisherOptions.rclPublisherOptions,
	)

	if ret != cwrap.Ok {
		return NewErr("RclInitOptionsInit", ret)
	}

	return nil
}

func (p *Publisher) Fini(node Node) error {
	ret := cwrap.RclPublisherFini(p.rclPublisher, node.rclNode)
	if ret != cwrap.Ok {
		return NewErr("RclPublisherFini", ret)
	}

	return nil
}

func (p *Publisher) Publish(msg RosMessage) error {

	ret := cwrap.RclPublish(
		p.rclPublisher,
		msg.Pointer(),
		nil,
	)

	if ret != cwrap.Ok {
		return NewErr("RclPublish", ret)
	}

	return nil
}

func (p *Publisher) IsValid() bool {
	return cwrap.RclPublisherIsValid(p.rclPublisher)
}
