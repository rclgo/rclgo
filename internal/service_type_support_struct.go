package cwrap

// #cgo CFLAGS: -I/opt/ros/eloquent/include
// #include <rcl/rcl.h>
import "C"

type RosidlServiceTypeSupport C.rosidl_service_type_support_t
