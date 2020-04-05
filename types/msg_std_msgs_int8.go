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

type StdMsgsInt8 struct {
	data    *cwrap.StdMsgs_MsgInt8
	MsgType MessageTypeSupport
}

func (msg *StdMsgsInt8) GetMessage() MessageTypeSupport {
	return msg.MsgType
}

func (msg *StdMsgsInt8) GetData() MessageData {
	return MessageData{unsafe.Pointer(msg.data)}
}

func (msg *StdMsgsInt8) InitMessage() {
	msg.data = cwrap.InitStdMsgsMsgInt8()
	msg.MsgType = GetMessageTypeFromStdMsgsInt8()
}

func (msg *StdMsgsInt8) SetInt8(data int8) {
	//TODO: to implement the setter
	msg.data.Set(data)
}

func (msg *StdMsgsInt8) GetInt8() int8 {
	return int8(msg.data.Data())
}

func (msg *StdMsgsInt8) GetDataAsString() string {
	//TODO: to implement the stringify opt
	myRetValue := ""
	return myRetValue
}

func (msg *StdMsgsInt8) DestroyMessage() {
	cwrap.DestroyStdMsgsMsgInt8(msg.data)
}

func GetMessageTypeFromStdMsgsInt8() MessageTypeSupport {
	ret := cwrap.GetMessageTypeFromStdMsgsMsgInt8()
	return MessageTypeSupport{ret}
}
