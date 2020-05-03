package cwrap

// #cgo CFLAGS: -I/opt/ros/eloquent/include
// #include <rcl/rcl.h>
import "C"

//
type RmwContext C.rmw_context_t

//
type RmwEvent C.rmw_event_t

//
type RmwGuardCondition C.rmw_guard_condition_t

//
type RmwPublisher C.rmw_publisher_t

//
type RmwPublisherAllocation C.rmw_publisher_allocation_t

//
type RmwQosProfile C.rmw_qos_profile_t

//
type RmwRet C.rmw_ret_t

//
type RmwService C.rmw_service_t

//
type RmwSubscription C.rmw_subscription_t
