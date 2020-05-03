package cwrap

// #cgo CFLAGS: -I/opt/ros/eloquent/include
// #include <rcl/rcl.h>
import "C"

type RclContextInstanceID C.rcl_context_instance_id_t

type RclContext C.rcl_context_t

// RclGetZeroInitializedContext returns a zero initialization context object.
func RclGetZeroInitializedContext() RclContext {
	var ret C.rcl_context_t = C.rcl_get_zero_initialized_context()
	return RclContext(ret)
}

// RclContextFini finalize a context.
func RclContextFini(
	context RclContextPtr,
) RclRet {
	var ret C.rcl_ret_t = C.rcl_context_fini(
		(*C.rcl_context_t)(context),
	)
	return RclRet(ret)
}

// RclContextGetInitOptions returns the init options used during initialization for this context.
func RclContextGetInitOptions(
	context *RclContext,
) *RclInitOptions {
	var ret *C.rcl_init_options_t = C.rcl_context_get_init_options(
		(*C.rcl_context_t)(context),
	)
	return (*RclInitOptions)(ret)
}

// RclContextGetInstanceID returns an unsigned integer that is unique to the given context, or `0` if invalid.
func RclContextGetInstanceID(
	context *RclContext,
) RclContextInstanceID {
	var ret C.rcl_context_instance_id_t = C.rcl_context_get_instance_id(
		(*C.rcl_context_t)(context),
	)
	return RclContextInstanceID(ret)
}

// RclContextIsValid returns `true` if the given context is currently valid, otherwise `false`.
func RclContextIsValid(
	context RclContextPtr,
) bool {
	var ret C.bool = C.rcl_context_is_valid(
		(*C.rcl_context_t)(context),
	)
	return bool(ret)
}

// RclContextGetRmwContext returns pointer to the rmw context if the given context is currently valid, otherwise `NULL`.
func RclContextGetRmwContext(
	context *RclContext,
) *RmwContext {
	var ret *C.rmw_context_t = C.rcl_context_get_rmw_context(
		(*C.rcl_context_t)(context),
	)
	return (*RmwContext)(ret)
}
