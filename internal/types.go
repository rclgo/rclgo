package cwrap

// #cgo CFLAGS: -I/opt/ros/eloquent/include
// #include <rcl/rcl.h>
import "C"

//
type RclSerializedMessage C.rcl_serialized_message_t
