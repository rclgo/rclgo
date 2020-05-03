package cwrap

// #cgo CFLAGS: -I/opt/ros/eloquent/include
// #include <rcl/rcl.h>
import "C"

import "unsafe"

// Structure which encapsulates a ROS QoS event handle.
type RclEvent C.rcl_event_t

// RclGetZeroInitializedEvent returns a rcl_event_t struct with members set to `NULL`.
func RclGetZeroInitializedEvent() RclEvent {
	var ret C.rcl_event_t = C.rcl_get_zero_initialized_event()
	return RclEvent(ret)
}

// RclPublisherEventInit initialize an rcl_event_t with a publisher.
func RclPublisherEventInit(
	event *RclEvent,
	publisher *RclPublisher,
	eventType RclPublisherEventType,
) RclRet {
	var ret C.rcl_ret_t = C.rcl_publisher_event_init(
		(*C.rcl_event_t)(event),
		(*C.rcl_publisher_t)(publisher),
		(C.rcl_publisher_event_type_t)(eventType),
	)
	return RclRet(ret)
}

// RclSubscriptionEventInit initialize an rcl_event_t with a subscription.
func RclSubscriptionEventInit(
	event *RclEvent,
	subscription *RclSubscription,
	eventType RclSubscriptionEventType,
) RclRet {
	var ret C.rcl_ret_t = C.rcl_subscription_event_init(
		(*C.rcl_event_t)(event),
		(*C.rcl_subscription_t)(subscription),
		(C.rcl_subscription_event_type_t)(eventType),
	)
	return RclRet(ret)
}

// RclTakeEvent take event using the event handle.
func RclTakeEvent(
	event *RclEvent,
	eventInfo unsafe.Pointer,
) RclRet {
	var ret C.rcl_ret_t = C.rcl_take_event(
		(*C.rcl_event_t)(event),
		eventInfo,
	)
	return RclRet(ret)
}

// RclEventFini finalize an event.
func RclEventFini(
	event *RclEvent,
) RclRet {
	var ret C.rcl_ret_t = C.rcl_event_fini(
		(*C.rcl_event_t)(event),
	)
	return RclRet(ret)
}

// RclEventGetRmwHandle returns the rmw event handle.
func RclEventGetRmwHandle(
	event *RclEvent,
) *RmwEvent {
	var ret *C.rmw_event_t = C.rcl_event_get_rmw_handle(
		(*C.rcl_event_t)(event),
	)
	return (*RmwEvent)(ret)
}
