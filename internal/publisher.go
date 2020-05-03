package cwrap

// #cgo CFLAGS: -I/opt/ros/eloquent/include
// #include <rcl/rcl.h>
import "C"

import "unsafe"

type RclPublisher C.rcl_publisher_t

// RclPublisherOptions are the options available for a rcl publisher.
type RclPublisherOptions C.rcl_publisher_options_t

// RclGetZeroInitializedPublisher returns a rcl_publisher_t struct with members
// set to `NULL`.
func RclGetZeroInitializedPublisher() RclPublisher {
	var ret C.rcl_publisher_t = C.rcl_get_zero_initialized_publisher()
	return RclPublisher(ret)
}

// RclPublisherInit initializes an rcl publisher.
func RclPublisherInit(
	publisher *RclPublisher,
	node *RclNode,
	typeSupport *RosidlMessageTypeSupport,
	topicName string,
	options *RclPublisherOptions,
) RclRet {
	tName := C.CString(topicName)
	defer C.free(unsafe.Pointer(tName))

	var ret C.rcl_ret_t = C.rcl_publisher_init(
		(*C.rcl_publisher_t)(publisher),
		(*C.rcl_node_t)(node),
		(*C.rosidl_message_type_support_t)(typeSupport),
		tName,
		(*C.rcl_publisher_options_t)(options),
	)
	return RclRet(ret)
}

// RclPublisherFini finalizes a rcl_publisher_t.
func RclPublisherFini(
	publisher *RclPublisher,
	node *RclNode,
) RclRet {
	var ret C.rcl_ret_t = C.rcl_publisher_fini(
		(*C.rcl_publisher_t)(publisher),
		(*C.rcl_node_t)(node),
	)
	return RclRet(ret)
}

// RclPublisherGetDefaultOptions returns the default publisher options in an
// rcl_publisher_options_t.
func RclPublisherGetDefaultOptions() RclPublisherOptions {
	var ret C.rcl_publisher_options_t = C.rcl_publisher_get_default_options()
	return RclPublisherOptions(ret)
}

// RclBorrowLoanedMessage borrow a loaned message.
func RclBorrowLoanedMessage(
	publisher *RclPublisher,
	typeSupport *RosidlMessageTypeSupport,
	rosMessage unsafe.Pointer,
) RclRet {
	var ret C.rcl_ret_t = C.rcl_borrow_loaned_message(
		(*C.rcl_publisher_t)(publisher),
		(*C.rosidl_message_type_support_t)(typeSupport),
		&rosMessage,
	)
	return RclRet(ret)
}

// RclReturnLoanedMessageFromPublisher returns a loaned message previously
// borrowed from a publisher.
func RclReturnLoanedMessageFromPublisher(
	publisher *RclPublisher,
	loanedMessage unsafe.Pointer,
) RclRet {
	var ret C.rcl_ret_t = C.rcl_return_loaned_message_from_publisher(
		(*C.rcl_publisher_t)(publisher),
		loanedMessage,
	)
	return RclRet(ret)
}

// RclPublish publish a ROS message on a topic using a publisher.
func RclPublish(
	publisher *RclPublisher,
	rosMessage unsafe.Pointer,
	allocation *RmwPublisherAllocation,
) RclRet {
	var ret C.rcl_ret_t = C.rcl_publish(
		(*C.rcl_publisher_t)(publisher),
		rosMessage,
		(*C.rmw_publisher_allocation_t)(allocation),
	)
	return RclRet(ret)
}

// RclPublishSerializedMessage publish a serialized message on a topic using a
// publisher.
func RclPublishSerializedMessage(
	publisher *RclPublisher,
	serializedMessage *RclSerializedMessage,
	allocation *RmwPublisherAllocation,
) RclRet {
	var ret C.rcl_ret_t = C.rcl_publish_serialized_message(
		(*C.rcl_publisher_t)(publisher),
		(*C.rcl_serialized_message_t)(serializedMessage),
		(*C.rmw_publisher_allocation_t)(allocation),
	)
	return RclRet(ret)
}

// RclPublishLoanedMessage publish a loaned message on a topic using a
// publisher.
func RclPublishLoanedMessage(
	publisher *RclPublisher,
	rosMessage unsafe.Pointer,
	allocation *RmwPublisherAllocation,
) RclRet {
	var ret C.rcl_ret_t = C.rcl_publish_loaned_message(
		(*C.rcl_publisher_t)(publisher),
		rosMessage,
		(*C.rmw_publisher_allocation_t)(allocation),
	)
	return RclRet(ret)
}

// RclPublisherAssertLiveliness manually assert that this Publisher is alive
// (for RMW_QOS_POLICY_LIVELINESS_MANUAL_BY_TOPIC)
func RclPublisherAssertLiveliness(
	publisher *RclPublisher,
) RclRet {
	var ret C.rcl_ret_t = C.rcl_publisher_assert_liveliness(
		(*C.rcl_publisher_t)(publisher),
	)
	return RclRet(ret)
}

// RclPublisherGetTopicName gets the topic name for the publisher.
func RclPublisherGetTopicName(
	publisher *RclPublisher,
) string {
	var ret *C.char = C.rcl_publisher_get_topic_name(
		(*C.rcl_publisher_t)(publisher),
	)
	return C.GoString(ret)
}

// RclPublisherGetOptions returns the rcl publisher options.
func RclPublisherGetOptions(
	publisher *RclPublisher,
) *RclPublisherOptions {
	var ret *C.rcl_publisher_options_t = C.rcl_publisher_get_options(
		(*C.rcl_publisher_t)(publisher),
	)
	return (*RclPublisherOptions)(ret)
}

// RclPublisherGetRmwHandle returns the rmw publisher handle.
func RclPublisherGetRmwHandle(
	publisher *RclPublisher,
) *RmwPublisher {
	var ret *C.rmw_publisher_t = C.rcl_publisher_get_rmw_handle(
		(*C.rcl_publisher_t)(publisher),
	)
	return (*RmwPublisher)(ret)
}

// RclPublisherGetContext returns the context associated with this publisher.
func RclPublisherGetContext(
	publisher *RclPublisher,
) *RclContext {
	var ret *C.rcl_context_t = C.rcl_publisher_get_context(
		(*C.rcl_publisher_t)(publisher),
	)
	return (*RclContext)(ret)
}

// RclPublisherIsValid returns true if the publisher is valid, otherwise false.
func RclPublisherIsValid(
	publisher *RclPublisher,
) bool {
	var ret C.bool = C.rcl_publisher_is_valid(
		(*C.rcl_publisher_t)(publisher),
	)
	return bool(ret)
}

// RclPublisherIsValidExceptContext returns true if the publisher is valid
// except the context, otherwise false.
func RclPublisherIsValidExceptContext(
	publisher *RclPublisher,
) bool {
	var ret C.bool = C.rcl_publisher_is_valid_except_context(
		(*C.rcl_publisher_t)(publisher),
	)
	return bool(ret)
}

// RclPublisherGetSubscriptionCount gets the number of subscriptions matched to
// a publisher.
func RclPublisherGetSubscriptionCount(
	publisher *RclPublisher,
	subscriptionCount *uint64,
) RmwRet {
	var ret C.rmw_ret_t = C.rcl_publisher_get_subscription_count(
		(*C.rcl_publisher_t)(publisher),
		(*C.size_t)(subscriptionCount),
	)
	return RmwRet(ret)
}

// RclPublisherGetActualQos gets the actual qos settings of the publisher.
func RclPublisherGetActualQos(
	publisher *RclPublisher,
) *RmwQosProfile {
	var ret *C.rmw_qos_profile_t = C.rcl_publisher_get_actual_qos(
		(*C.rcl_publisher_t)(publisher),
	)
	return (*RmwQosProfile)(ret)
}

// RclPublisherCanLoanMessages checks if publisher instance can loan messages.
func RclPublisherCanLoanMessages(
	publisher *RclPublisher,
) bool {
	var ret C.bool = C.rcl_publisher_can_loan_messages(
		(*C.rcl_publisher_t)(publisher),
	)
	return bool(ret)
}
