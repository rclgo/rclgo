package cwrap

// #cgo CFLAGS: -I/opt/ros/eloquent/include
// #include <rcl/rcl.h>
import "C"

//
type RosidlMessageTypeSupport C.rosidl_message_type_support_t
