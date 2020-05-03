package cwrap

// #cgo CFLAGS: -I/opt/ros/eloquent/include
// #include <rcl/rcl.h>
import "C"

// Container for subscription's, guard condition's, etc to be waited on.
type RclWaitSet C.rcl_wait_set_t

// RclGetZeroInitializedWaitSet returns a rcl_wait_set_t struct with members set to `NULL`.
func RclGetZeroInitializedWaitSet() RclWaitSet {
	var ret C.rcl_wait_set_t = C.rcl_get_zero_initialized_wait_set()
	return RclWaitSet(ret)
}

// RclWaitSetInit initialize an rcl wait set with space for items to be waited on.
func RclWaitSetInit(
	waitSet *RclWaitSet,
	numberOfSubscriptions uint,
	numberOfGuardConditions uint,
	numberOfTimers uint,
	numberOfClients uint,
	numberOfServices uint,
	numberOfEvents uint,
	context RclContextPtr,
	allocator RclAllocator,
) RclRet {
	var ret C.rcl_ret_t = C.rcl_wait_set_init(
		(*C.rcl_wait_set_t)(waitSet),
		(C.size_t)(numberOfSubscriptions),
		(C.size_t)(numberOfGuardConditions),
		(C.size_t)(numberOfTimers),
		(C.size_t)(numberOfClients),
		(C.size_t)(numberOfServices),
		(C.size_t)(numberOfEvents),
		(*C.rcl_context_t)(context),
		(C.rcl_allocator_t)(allocator),
	)
	return RclRet(ret)
}

// RclWaitSetFini finalize an rcl wait set.
func RclWaitSetFini(
	waitSet *RclWaitSet,
) RclRet {
	var ret C.rcl_ret_t = C.rcl_wait_set_fini(
		(*C.rcl_wait_set_t)(waitSet),
	)
	return RclRet(ret)
}

// RclWaitSetGetAllocator retrieve the wait set's allocator.
func RclWaitSetGetAllocator(
	waitSet *RclWaitSet,
	allocator *RclAllocator,
) RclRet {
	var ret C.rcl_ret_t = C.rcl_wait_set_get_allocator(
		(*C.rcl_wait_set_t)(waitSet),
		(*C.rcl_allocator_t)(allocator),
	)
	return RclRet(ret)
}

// RclWaitSetAddSubscription store a pointer to the given subscription in the next empty spot in the set.
func RclWaitSetAddSubscription(
	waitSet *RclWaitSet,
	subscription *RclSubscription,
	index *uint64,
) RclRet {
	var ret C.rcl_ret_t = C.rcl_wait_set_add_subscription(
		(*C.rcl_wait_set_t)(waitSet),
		(*C.rcl_subscription_t)(subscription),
		(*C.size_t)(index),
	)
	return RclRet(ret)
}

// RclWaitSetClear remove (sets to `NULL`) all entities in the wait set.
func RclWaitSetClear(
	waitSet *RclWaitSet,
) RclRet {
	var ret C.rcl_ret_t = C.rcl_wait_set_clear(
		(*C.rcl_wait_set_t)(waitSet),
	)
	return RclRet(ret)
}

// RclWaitSetResize reallocate space for entities in the wait set.
func RclWaitSetResize(
	waitSet *RclWaitSet,
	subscriptionsSize uint,
	guardConditionsSize uint,
	timersSize uint,
	clientsSize uint,
	servicesSize uint,
	eventsSize uint,
) RclRet {
	var ret C.rcl_ret_t = C.rcl_wait_set_resize(
		(*C.rcl_wait_set_t)(waitSet),
		(C.size_t)(subscriptionsSize),
		(C.size_t)(guardConditionsSize),
		(C.size_t)(timersSize),
		(C.size_t)(clientsSize),
		(C.size_t)(servicesSize),
		(C.size_t)(eventsSize),
	)
	return RclRet(ret)
}

// RclWaitSetAddGuardCondition store a pointer to the guard condition in the next empty spot in the set.
func RclWaitSetAddGuardCondition(
	waitSet *RclWaitSet,
	guardCondition *RclGuardCondition,
	index *uint64,
) RclRet {
	var ret C.rcl_ret_t = C.rcl_wait_set_add_guard_condition(
		(*C.rcl_wait_set_t)(waitSet),
		(*C.rcl_guard_condition_t)(guardCondition),
		(*C.size_t)(index),
	)
	return RclRet(ret)
}

// RclWaitSetAddTimer store a pointer to the timer in the next empty spot in the set.
func RclWaitSetAddTimer(
	waitSet *RclWaitSet,
	timer *RclTimer,
	index *uint64,
) RclRet {
	var ret C.rcl_ret_t = C.rcl_wait_set_add_timer(
		(*C.rcl_wait_set_t)(waitSet),
		(*C.rcl_timer_t)(timer),
		(*C.size_t)(index),
	)
	return RclRet(ret)
}

// RclWaitSetAddClient store a pointer to the client in the next empty spot in the set.
func RclWaitSetAddClient(
	waitSet *RclWaitSet,
	client *RclClient,
	index *uint64,
) RclRet {
	var ret C.rcl_ret_t = C.rcl_wait_set_add_client(
		(*C.rcl_wait_set_t)(waitSet),
		(*C.rcl_client_t)(client),
		(*C.size_t)(index),
	)
	return RclRet(ret)
}

// RclWaitSetAddService store a pointer to the service in the next empty spot in the set.
func RclWaitSetAddService(
	waitSet *RclWaitSet,
	service *RclService,
	index *uint64,
) RclRet {
	var ret C.rcl_ret_t = C.rcl_wait_set_add_service(
		(*C.rcl_wait_set_t)(waitSet),
		(*C.rcl_service_t)(service),
		(*C.size_t)(index),
	)
	return RclRet(ret)
}

// RclWaitSetAddEvent store a pointer to the event in the next empty spot in the set.
func RclWaitSetAddEvent(
	waitSet *RclWaitSet,
	event *RclEvent,
	index *uint64,
) RclRet {
	var ret C.rcl_ret_t = C.rcl_wait_set_add_event(
		(*C.rcl_wait_set_t)(waitSet),
		(*C.rcl_event_t)(event),
		(*C.size_t)(index),
	)
	return RclRet(ret)
}

// RclWait block until the wait set is ready or until the timeout has been exceeded.
func RclWait(
	waitSet *RclWaitSet,
	timeout int64,
) RclRet {
	var ret C.rcl_ret_t = C.rcl_wait(
		(*C.rcl_wait_set_t)(waitSet),
		(C.int64_t)(timeout),
	)
	return RclRet(ret)
}
