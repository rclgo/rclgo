package cwrap

// #cgo CFLAGS: -I/opt/ros/eloquent/include
// #include <rcl/rcl.h>
import "C"

import (
	"unsafe"
)

type RclSubscription C.rcl_subscription_t

// RclSubscriptionOptions are the options available for a rcl subscription.
type RclSubscriptionOptions C.rcl_subscription_options_t

// RclGetZeroInitializedSubscription returns a rcl_subscription_t struct with
// members set to `NULL`.
func RclGetZeroInitializedSubscription() RclSubscription {
	var ret C.rcl_subscription_t = C.rcl_get_zero_initialized_subscription()
	return RclSubscription(ret)
}

// RclSubscriptionInit initializes a ROS subscription.
func RclSubscriptionInit(
	subscription *RclSubscription,
	node *RclNode,
	typeSupport *RosidlMessageTypeSupport,
	topicName string,
	options *RclSubscriptionOptions,
) RclRet {
	cTopic := C.CString(topicName)
	defer C.free(unsafe.Pointer(cTopic))
	var ret C.rcl_ret_t = C.rcl_subscription_init(
		(*C.rcl_subscription_t)(subscription),
		(*C.rcl_node_t)(node),
		(*C.rosidl_message_type_support_t)(typeSupport),
		cTopic,
		(*C.rcl_subscription_options_t)(options),
	)
	return RclRet(ret)
}

// RclSubscriptionFini finalizes a rcl_subscription_t.
func RclSubscriptionFini(
	subscription *RclSubscription,
	node *RclNode,
) RclRet {
	var ret C.rcl_ret_t = C.rcl_subscription_fini(
		(*C.rcl_subscription_t)(subscription),
		(*C.rcl_node_t)(node),
	)
	return RclRet(ret)
}

// RclSubscriptionGetDefaultOptions returns the default subscription options in
// an rcl_subscription_options_t.
func RclSubscriptionGetDefaultOptions() RclSubscriptionOptions {
	var ret C.rcl_subscription_options_t = C.rcl_subscription_get_default_options()
	return RclSubscriptionOptions(ret)
}

// RclTake takes a ROS message from a topic using an rcl subscription.
func RclTake(
	subscription *RclSubscription,
	rosMessage unsafe.Pointer,
	messageInfo *RmwMessageInfo,
	allocation *RmwSubscriptionAllocation,
) RclRet {
	var ret C.rcl_ret_t = C.rcl_take(
		(*C.rcl_subscription_t)(subscription),
		rosMessage,
		(*C.rmw_message_info_t)(messageInfo),
		(*C.rmw_subscription_allocation_t)(allocation),
	)
	return RclRet(ret)
}

// RclTakeSerializedMessage takes a serialized raw message from a topic using an
//  rcl subscription.
func RclTakeSerializedMessage(
	subscription *RclSubscription,
	serializedMessage *RclSerializedMessage,
	messageInfo *RmwMessageInfo,
	allocation *RmwSubscriptionAllocation,
) RclRet {
	var ret C.rcl_ret_t = C.rcl_take_serialized_message(
		(*C.rcl_subscription_t)(subscription),
		(*C.rcl_serialized_message_t)(serializedMessage),
		(*C.rmw_message_info_t)(messageInfo),
		(*C.rmw_subscription_allocation_t)(allocation),
	)
	return RclRet(ret)
}

// RclTakeLoanedMessage takes a loaned message from a topic using an rcl
// subscription.
func RclTakeLoanedMessage(
	subscription *RclSubscription,
	loanedMessage *unsafe.Pointer,
	messageInfo *RmwMessageInfo,
	allocation *RmwSubscriptionAllocation,
) RclRet {
	var ret C.rcl_ret_t = C.rcl_take_loaned_message(
		(*C.rcl_subscription_t)(subscription),
		loanedMessage,
		(*C.rmw_message_info_t)(messageInfo),
		(*C.rmw_subscription_allocation_t)(allocation),
	)
	return RclRet(ret)
}

// RclReturnLoanedMessageFromSubscription returns a loaned message from a topic
// using an rcl subscription.
func RclReturnLoanedMessageFromSubscription(
	subscription *RclSubscription,
	loanedMessage unsafe.Pointer,
) RclRet {
	var ret C.rcl_ret_t = C.rcl_return_loaned_message_from_subscription(
		(*C.rcl_subscription_t)(subscription),
		loanedMessage,
	)
	return RclRet(ret)
}

// RclSubscriptionGetTopicName gets the topic name for the subscription.
func RclSubscriptionGetTopicName(
	subscription *RclSubscription,
) string {
	var ret *C.char = C.rcl_subscription_get_topic_name(
		(*C.rcl_subscription_t)(subscription),
	)
	return C.GoString(ret)
}

// RclSubscriptionGetOptions returns the rcl subscription options.
func RclSubscriptionGetOptions(
	subscription *RclSubscription,
) *RclSubscriptionOptions {
	var ret *C.rcl_subscription_options_t = C.rcl_subscription_get_options(
		(*C.rcl_subscription_t)(subscription),
	)
	return (*RclSubscriptionOptions)(ret)
}

// RclSubscriptionGetRmwHandle returns the rmw subscription handle.
func RclSubscriptionGetRmwHandle(
	subscription *RclSubscription,
) *RmwSubscription {
	var ret *C.rmw_subscription_t = C.rcl_subscription_get_rmw_handle(
		(*C.rcl_subscription_t)(subscription),
	)
	return (*RmwSubscription)(ret)
}

// RclSubscriptionIsValid checks that the subscription is valid.
func RclSubscriptionIsValid(
	subscription *RclSubscription,
) bool {
	var ret C.bool = C.rcl_subscription_is_valid(
		(*C.rcl_subscription_t)(subscription),
	)
	return bool(ret)
}

// RclSubscriptionGetPublisherCount gets the number of publishers matched to a
// subscription.
func RclSubscriptionGetPublisherCount(
	subscription *RclSubscription,
	publisherCount *uint64,
) RmwRet {
	var ret C.rmw_ret_t = C.rcl_subscription_get_publisher_count(
		(*C.rcl_subscription_t)(subscription),
		(*C.size_t)(publisherCount),
	)
	return RmwRet(ret)
}

// RclSubscriptionGetActualQos gets the actual qos settings of the subscription.
func RclSubscriptionGetActualQos(
	subscription *RclSubscription,
) *RmwQosProfile {
	var ret *C.rmw_qos_profile_t = C.rcl_subscription_get_actual_qos(
		(*C.rcl_subscription_t)(subscription),
	)
	return (*RmwQosProfile)(ret)
}

// RclSubscriptionCanLoanMessages checks if subscription instance can loan
// messages.
func RclSubscriptionCanLoanMessages(
	subscription *RclSubscription,
) bool {
	var ret C.bool = C.rcl_subscription_can_loan_messages(
		(*C.rcl_subscription_t)(subscription),
	)
	return bool(ret)
}
