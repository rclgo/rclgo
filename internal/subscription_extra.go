package cwrap

// #cgo CFLAGS: -I/opt/ros/eloquent/include
// #include <rcl/rcl.h>
import "C"

type RmwSubscriptionAllocation C.rmw_subscription_allocation_t
