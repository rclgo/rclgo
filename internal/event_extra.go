package cwrap

// #cgo CFLAGS: -I/opt/ros/eloquent/include
// #include <rcl/rcl.h>
import "C"

//
type RclPublisherEventType C.rcl_publisher_event_type_t

const (
	RclPublisherOfferedDeadlineMissed = iota
	RclPublisherLivelinessLost
)

//
type RclSubscriptionEventType C.rcl_subscription_event_type_t

const (
	RclSubscriptionRequestdDeadlineMissed = iota
	RclSubscriptionLivelinessChanged
)
