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

type StdMsgsInt16 struct {
	data    *cwrap.StdMsgs_MsgInt16
	MsgType MessageTypeSupport
}

func (msg *StdMsgsInt16) GetMessage() MessageTypeSupport {
	return msg.MsgType
}

func (msg *StdMsgsInt16) GetData() MessageData {
	return MessageData{unsafe.Pointer(msg.data)}
}

func (msg *StdMsgsInt16) InitMessage() {
	msg.data = cwrap.InitStdMsgsMsgInt16()
	msg.MsgType = GetMessageTypeFromStdMsgsInt16()
}

func (msg *StdMsgsInt16) SetInt16(data int16) {
	//TODO: to implement the setter
	msg.data.Set(data)
}

func (msg *StdMsgsInt16) GetInt16() int16 {
	return int16(msg.data.Data())
}

func (msg *StdMsgsInt16) GetDataAsString() string {
	//TODO: to implement the stringify opt
	myRetValue := ""
	return myRetValue
}

func (msg *StdMsgsInt16) DestroyMessage() {
	cwrap.DestroyStdMsgsMsgInt16(msg.data)
}

func GetMessageTypeFromStdMsgsInt16() MessageTypeSupport {
	ret := cwrap.GetMessageTypeFromStdMsgsMsgInt16()
	return MessageTypeSupport{ret}
}
