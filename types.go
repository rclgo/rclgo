package rclgo

import (
	"unsafe"

	cwrap "github.com/rclgo/rclgo/internal"
)

//
type MessageTypeSupport struct {
	RosidlMessageTypeSupport *cwrap.RosidlMessageTypeSupport
}

// NewMessageTypeSupport returns a rosidl_message_type_support_t wrapper. rmts
// MUST be cast-able to a rosidl_message_type_support_t*
func NewMessageTypeSupport(rmts unsafe.Pointer) MessageTypeSupport {
	inner := (*cwrap.RosidlMessageTypeSupport)(rmts)
	// if !ok {
	// 	panic("interface could not be asserted as RosidlMessageTypeSupport")
	// }
	return MessageTypeSupport{RosidlMessageTypeSupport: inner}
}

type MessageInfo struct {
	RmwMessageInfo *cwrap.RmwMessageInfo
}

//
type MessageData struct {
	Data unsafe.Pointer
}

//
type Message interface {
	MessageType() MessageTypeSupport
	Data() MessageData
	Init() error
	Destroy() error
}
