package cwrap

// #cgo CFLAGS: -I/opt/ros/eloquent/include
// #include <rcl/rcl.h>
import "C"

// Structure which encapsulates a ROS Timer.
type RclTimer C.rcl_timer_t

type RclTimerCallback C.rcl_timer_callback_t

// RclGetZeroInitializedTimer returns a zero initialized timer.
func RclGetZeroInitializedTimer() RclTimer {
	var ret C.rcl_timer_t = C.rcl_get_zero_initialized_timer()
	return RclTimer(ret)
}

// RclTimerInit initialize a timer.
func RclTimerInit(
	timer *RclTimer,
	clock *RclClock,
	context *RclContext,
	period int64,
	callback RclTimerCallback,
	allocator RclAllocator,
) RclRet {
	var ret C.rcl_ret_t = C.rcl_timer_init(
		(*C.rcl_timer_t)(timer),
		(*C.rcl_clock_t)(clock),
		(*C.rcl_context_t)(context),
		(C.int64_t)(period),
		(C.rcl_timer_callback_t)(callback),
		(C.rcl_allocator_t)(allocator),
	)
	return RclRet(ret)
}

// RclTimerFini finalize a timer.
func RclTimerFini(
	timer *RclTimer,
) RclRet {
	var ret C.rcl_ret_t = C.rcl_timer_fini(
		(*C.rcl_timer_t)(timer),
	)
	return RclRet(ret)
}

// RclTimerCall call the timer's callback and set the last call time.
func RclTimerCall(
	timer *RclTimer,
) RclRet {
	var ret C.rcl_ret_t = C.rcl_timer_call(
		(*C.rcl_timer_t)(timer),
	)
	return RclRet(ret)
}

// RclTimerClock retrieve the clock of the timer.
func RclTimerClock(
	timer *RclTimer,
	clock **RclClock,
) RclRet {
	var cclock = (*C.rcl_clock_t)(*clock)
	var ret C.rcl_ret_t = C.rcl_timer_clock(
		(*C.rcl_timer_t)(timer),
		// ToDo: Check this
		&cclock,
	)
	return RclRet(ret)
}

// RclTimerIsReady calculates whether or not the timer should be called.
func RclTimerIsReady(
	timer *RclTimer,
	isReady *bool,
) RclRet {
	var ret C.rcl_ret_t = C.rcl_timer_is_ready(
		(*C.rcl_timer_t)(timer),
		(*C.bool)(isReady),
	)
	return RclRet(ret)
}

// RclTimerGetTimeUntilNextCall calculate and retrieve the time until the next call in nanoseconds.
func RclTimerGetTimeUntilNextCall(
	timer *RclTimer,
	timeUntilNextCall *int64,
) RclRet {
	var ret C.rcl_ret_t = C.rcl_timer_get_time_until_next_call(
		(*C.rcl_timer_t)(timer),
		(*C.int64_t)(timeUntilNextCall),
	)
	return RclRet(ret)
}

// RclTimerGetTimeSinceLastCall retrieve the time since the previous call to rcl_timer_call() occurred.
func RclTimerGetTimeSinceLastCall(
	timer *RclTimer,
	timeSinceLastCall *int64,
) RclRet {
	var ret C.rcl_ret_t = C.rcl_timer_get_time_since_last_call(
		(*C.rcl_timer_t)(timer),
		(*C.int64_t)(timeSinceLastCall),
	)
	return RclRet(ret)
}

// RclTimerGetPeriod retrieve the period of the timer.
func RclTimerGetPeriod(
	timer *RclTimer,
	period *int64,
) RclRet {
	var ret C.rcl_ret_t = C.rcl_timer_get_period(
		(*C.rcl_timer_t)(timer),
		(*C.int64_t)(period),
	)
	return RclRet(ret)
}

// RclTimerExchangePeriod exchange the period of the timer and return the previous period.
func RclTimerExchangePeriod(
	timer *RclTimer,
	newPeriod int64,
	oldPeriod *int64,
) RclRet {
	var ret C.rcl_ret_t = C.rcl_timer_exchange_period(
		(*C.rcl_timer_t)(timer),
		(C.int64_t)(newPeriod),
		(*C.int64_t)(oldPeriod),
	)
	return RclRet(ret)
}

// RclTimerGetCallback returns the current timer callback.
func RclTimerGetCallback(
	timer *RclTimer,
) RclTimerCallback {
	var ret C.rcl_timer_callback_t = C.rcl_timer_get_callback(
		(*C.rcl_timer_t)(timer),
	)
	return RclTimerCallback(ret)
}

// RclTimerExchangeCallback exchange the current timer callback and return the current callback.
func RclTimerExchangeCallback(
	timer *RclTimer,
	newCallback RclTimerCallback,
) RclTimerCallback {
	var ret C.rcl_timer_callback_t = C.rcl_timer_exchange_callback(
		(*C.rcl_timer_t)(timer),
		(C.rcl_timer_callback_t)(newCallback),
	)
	return RclTimerCallback(ret)
}

// RclTimerCancel cancel a timer.
func RclTimerCancel(
	timer *RclTimer,
) RclRet {
	var ret C.rcl_ret_t = C.rcl_timer_cancel(
		(*C.rcl_timer_t)(timer),
	)
	return RclRet(ret)
}

// RclTimerIsCanceled retrieve the canceled state of a timer.
func RclTimerIsCanceled(
	timer *RclTimer,
	isCanceled *bool,
) RclRet {
	var ret C.rcl_ret_t = C.rcl_timer_is_canceled(
		(*C.rcl_timer_t)(timer),
		(*C.bool)(isCanceled),
	)
	return RclRet(ret)
}

// RclTimerReset reset a timer.
func RclTimerReset(
	timer *RclTimer,
) RclRet {
	var ret C.rcl_ret_t = C.rcl_timer_reset(
		(*C.rcl_timer_t)(timer),
	)
	return RclRet(ret)
}

// RclTimerGetGuardCondition retrieve a guard condition used by the timer to wake the waitset when using ROSTime.
func RclTimerGetGuardCondition(
	timer *RclTimer,
) *RclGuardCondition {
	var ret *C.rcl_guard_condition_t = C.rcl_timer_get_guard_condition(
		(*C.rcl_timer_t)(timer),
	)
	return (*RclGuardCondition)(ret)
}
