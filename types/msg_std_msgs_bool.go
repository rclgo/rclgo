package types

/////////////////////////////////////////////////////
//// THE CONTENT OF THIS FILE HAS BEEN AUTOGENERATED
/////////////////////////////////////////////////////
// #cgo CFLAGS: -I/opt/ros/eloquent/include
// #cgo LDFLAGS: -L/opt/ros/eloquent/lib -Wl,-rpath=/opt/ros/eloquent/lib -lrcl -lrosidl_generator_c -lrosidl_typesupport_c -lstd_msgs__rosidl_generator_c -lstd_msgs__rosidl_typesupport_c
// #include <rcl/rcl.h>
import "C"

import (
	"unsafe"

	"github.com/richardrigby/rclgo/cwrap"
)

type StdMsgsBool struct {
	data    *cwrap.StdMsgs_MsgBool
	MsgType MessageTypeSupport
}

func (msg *StdMsgsBool) GetMessage() MessageTypeSupport {
	return msg.MsgType
}

func (msg *StdMsgsBool) GetData() MessageData {
	return MessageData{unsafe.Pointer(msg.data)}
}

func (msg *StdMsgsBool) InitMessage() {
	msg.data = cwrap.InitStdMsgsMsgBool()
	msg.MsgType = GetMessageTypeFromStdMsgsBool()
}

func (msg *StdMsgsBool) SetBool(data bool) {
	//TODO: to implement the setter
	msg.data.Set(data)
}

func (msg *StdMsgsBool) GetBool() bool {
	return bool(msg.data.Data())
}

func (msg *StdMsgsBool) GetDataAsString() string {
	//TODO: to implement the stringify opt
	myRetValue := ""
	return myRetValue
}

func (msg *StdMsgsBool) DestroyMessage() {
	cwrap.DestroyStdMsgsMsgBool(msg.data)
}

func GetMessageTypeFromStdMsgsBool() MessageTypeSupport {
	ret := cwrap.GetMessageTypeFromStdMsgsMsgBool()
	return MessageTypeSupport{ret}
}
