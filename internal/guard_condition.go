package cwrap

// #cgo CFLAGS: -I/opt/ros/eloquent/include
// #include <rcl/rcl.h>
import "C"

// Handle for a rcl guard condition.
type RclGuardCondition C.rcl_guard_condition_t

// Options available for a rcl guard condition.
type RclGuardConditionOptions C.rcl_guard_condition_options_t

// RclGetZeroInitializedGuardCondition returns a rcl_guard_condition_t struct with members set to `NULL`.
func RclGetZeroInitializedGuardCondition() RclGuardCondition {
	var ret C.rcl_guard_condition_t = C.rcl_get_zero_initialized_guard_condition()
	return RclGuardCondition(ret)
}

// RclGuardConditionInit initialize an rcl guard_condition.
func RclGuardConditionInit(
	guardCondition *RclGuardCondition,
	context RclContextPtr,
	options RclGuardConditionOptions,
) RclRet {
	var ret C.rcl_ret_t = C.rcl_guard_condition_init(
		(*C.rcl_guard_condition_t)(guardCondition),
		(*C.rcl_context_t)(context),
		(C.rcl_guard_condition_options_t)(options),
	)
	return RclRet(ret)
}

// RclGuardConditionFini finalize a rcl_guard_condition_t.
func RclGuardConditionFini(
	guardCondition *RclGuardCondition,
) RclRet {
	var ret C.rcl_ret_t = C.rcl_guard_condition_fini(
		(*C.rcl_guard_condition_t)(guardCondition),
	)
	return RclRet(ret)
}

// RclGuardConditionGetDefaultOptions returns the default options in a rcl_guard_condition_options_t struct.
func RclGuardConditionGetDefaultOptions() RclGuardConditionOptions {
	var ret C.rcl_guard_condition_options_t = C.rcl_guard_condition_get_default_options()
	return RclGuardConditionOptions(ret)
}

// RclTriggerGuardCondition trigger an rcl guard condition.
func RclTriggerGuardCondition(
	guardCondition *RclGuardCondition,
) RclRet {
	var ret C.rcl_ret_t = C.rcl_trigger_guard_condition(
		(*C.rcl_guard_condition_t)(guardCondition),
	)
	return RclRet(ret)
}

// RclGuardConditionGetOptions returns the guard condition options.
func RclGuardConditionGetOptions(
	guardCondition *RclGuardCondition,
) *RclGuardConditionOptions {
	var ret *C.rcl_guard_condition_options_t = C.rcl_guard_condition_get_options(
		(*C.rcl_guard_condition_t)(guardCondition),
	)
	return (*RclGuardConditionOptions)(ret)
}

// RclGuardConditionGetRmwHandle returns the rmw guard condition handle.
func RclGuardConditionGetRmwHandle(
	guardCondition *RclGuardCondition,
) *RmwGuardCondition {
	var ret *C.rmw_guard_condition_t = C.rcl_guard_condition_get_rmw_handle(
		(*C.rcl_guard_condition_t)(guardCondition),
	)
	return (*RmwGuardCondition)(ret)
}
