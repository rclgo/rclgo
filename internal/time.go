package cwrap

// #cgo CFLAGS: -I/opt/ros/eloquent/include
// #include <rcl/rcl.h>
import "C"

type RclClockType uint32

const (
	RclClockUnitialized RclClockType = 1
	RclRosTime                       = 2
	RclSystemTime                    = 3
	RclSteadyTime                    = 4
)

type RclTimePointValue C.int64_t

// A duration of time, measured in nanoseconds and its source.
type RclDuration C.rcl_duration_t

// Struct to describe a jump in time.
type RclTimeJump C.rcl_time_jump_t

// Describe the prerequisites for calling a time jump callback.
type RclJumpThreshold C.rcl_jump_threshold_t

// Struct to describe an added callback.
type RclJumpCallbackInfo C.rcl_jump_callback_info_t

// Encapsulation of a time source.
type RclClock C.rcl_clock_t

// A single point in time, measured in nanoseconds, the reference point is based on the source.
type RclTimePoint C.rcl_time_point_t

// RclClockValid check if the clock has valid values.
func RclClockValid(
	clock *RclClock,
) bool {
	var ret C.bool = C.rcl_clock_valid(
		(*C.rcl_clock_t)(clock),
	)
	return bool(ret)
}

// RclClockInit initialize a clock based on the passed type.
func RclClockInit(
	clockType RclClockType,
	clock *RclClock,
	allocator *RclAllocator,
) RclRet {
	var ret C.rcl_ret_t = C.rcl_clock_init(
		uint32(clockType),
		(*C.rcl_clock_t)(clock),
		(*C.rcl_allocator_t)(allocator),
	)
	return RclRet(ret)
}

// RclClockFini finalize a clock.
func RclClockFini(
	clock *RclClock,
) RclRet {
	var ret C.rcl_ret_t = C.rcl_clock_fini(
		(*C.rcl_clock_t)(clock),
	)
	return RclRet(ret)
}

// RclRosClockInit initialize a clock as a RCL_ROS_TIME time source.
func RclRosClockInit(
	clock *RclClock,
	allocator *RclAllocator,
) RclRet {
	var ret C.rcl_ret_t = C.rcl_ros_clock_init(
		(*C.rcl_clock_t)(clock),
		(*C.rcl_allocator_t)(allocator),
	)
	return RclRet(ret)
}

// RclRosClockFini finalize a clock as a `RCL_ROS_TIME` time source.
func RclRosClockFini(
	clock *RclClock,
) RclRet {
	var ret C.rcl_ret_t = C.rcl_ros_clock_fini(
		(*C.rcl_clock_t)(clock),
	)
	return RclRet(ret)
}

// RclSteadyClockInit initialize a clock as a `RCL_STEADY_TIME` time source.
func RclSteadyClockInit(
	clock *RclClock,
	allocator *RclAllocator,
) RclRet {
	var ret C.rcl_ret_t = C.rcl_steady_clock_init(
		(*C.rcl_clock_t)(clock),
		(*C.rcl_allocator_t)(allocator),
	)
	return RclRet(ret)
}

// RclSteadyClockFini finalize a clock as a `RCL_STEADY_TIME` time source.
func RclSteadyClockFini(
	clock *RclClock,
) RclRet {
	var ret C.rcl_ret_t = C.rcl_steady_clock_fini(
		(*C.rcl_clock_t)(clock),
	)
	return RclRet(ret)
}

// RclSystemClockInit initialize a clock as a `RCL_SYSTEM_TIME` time source.
func RclSystemClockInit(
	clock *RclClock,
	allocator *RclAllocator,
) RclRet {
	var ret C.rcl_ret_t = C.rcl_system_clock_init(
		(*C.rcl_clock_t)(clock),
		(*C.rcl_allocator_t)(allocator),
	)
	return RclRet(ret)
}

// RclSystemClockFini finalize a clock as a `RCL_SYSTEM_TIME` time source.
func RclSystemClockFini(
	clock *RclClock,
) RclRet {
	var ret C.rcl_ret_t = C.rcl_system_clock_fini(
		(*C.rcl_clock_t)(clock),
	)
	return RclRet(ret)
}

// RclDifferenceTimes compute the difference between two time points
func RclDifferenceTimes(
	start *RclTimePoint,
	finish *RclTimePoint,
	delta *RclDuration,
) RclRet {
	var ret C.rcl_ret_t = C.rcl_difference_times(
		(*C.rcl_time_point_t)(start),
		(*C.rcl_time_point_t)(finish),
		(*C.rcl_duration_t)(delta),
	)
	return RclRet(ret)
}

// RclClockGetNow fill the time point value with the current value of the associated clock.
func RclClockGetNow(
	clock *RclClock,
	timePointValue *RclTimePointValue,
) RclRet {
	var ret C.rcl_ret_t = C.rcl_clock_get_now(
		(*C.rcl_clock_t)(clock),
		(*C.rcl_time_point_value_t)(timePointValue),
	)
	return RclRet(ret)
}

// RclEnableRosTimeOverride enable the ROS time abstraction override.
func RclEnableRosTimeOverride(
	clock *RclClock,
) RclRet {
	var ret C.rcl_ret_t = C.rcl_enable_ros_time_override(
		(*C.rcl_clock_t)(clock),
	)
	return RclRet(ret)
}

// RclDisableRosTimeOverride disable the ROS time abstraction override.
func RclDisableRosTimeOverride(
	clock *RclClock,
) RclRet {
	var ret C.rcl_ret_t = C.rcl_disable_ros_time_override(
		(*C.rcl_clock_t)(clock),
	)
	return RclRet(ret)
}

// RclIsEnabledRosTimeOverride check if the `RCL_ROS_TIME` time source has the override enabled.
func RclIsEnabledRosTimeOverride(
	clock *RclClock,
	isEnabled *bool,
) RclRet {
	var ret C.rcl_ret_t = C.rcl_is_enabled_ros_time_override(
		(*C.rcl_clock_t)(clock),
		(*C.bool)(isEnabled),
	)
	return RclRet(ret)
}

// RclSetRosTimeOverride set the current time for this `RCL_ROS_TIME` time source.
func RclSetRosTimeOverride(
	clock *RclClock,
	timeValue RclTimePointValue,
) RclRet {
	var ret C.rcl_ret_t = C.rcl_set_ros_time_override(
		(*C.rcl_clock_t)(clock),
		(C.rcl_time_point_value_t)(timeValue),
	)
	return RclRet(ret)
}

// // RclClockAddJumpCallback add a callback to be called when a time jump exceeds a threshold.
// func RclClockAddJumpCallback(
// 	clock *RclClock,
// 	threshold RclJumpThreshold,
// 	callback RclJumpCallback,
// 	userData unsafe.Pointer,
// ) RclRet {
// 	var ret C.rcl_ret_t = C.rcl_clock_add_jump_callback(
// 		(*C.rcl_clock_t)(clock),
// 		(C.rcl_jump_threshold_t)(threshold),
// 		(C.rcl_jump_callback_t)(callback),
// 		(*C.void)(userData),
// 	)
// 	return RclRet(ret)
// }

// // RclClockRemoveJumpCallback remove a previously added time jump callback.
// func RclClockRemoveJumpCallback(
// 	clock *RclClock,
// 	callback RclJumpCallback,
// 	userData unsafe.Pointer,
// ) RclRet {
// 	var ret C.rcl_ret_t = C.rcl_clock_remove_jump_callback(
// 		(*C.rcl_clock_t)(clock),
// 		(C.rcl_jump_callback_t)(callback),
// 		(*C.void)(userData),
// 	)
// 	return RclRet(ret)
// }
