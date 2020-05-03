package cwrap

// #cgo CFLAGS: -I/opt/ros/eloquent/include
// #include <rcl/rcl.h>
// #include <stdlib.h>
// void free_context(rcl_context_t *ctx) {
// 	free(ctx);
// }
import "C"

import "unsafe"

//
type RclInitOptions C.rcl_init_options_t

//
type RclAllocator C.rcl_allocator_t

//
func RclGetZeroInitializedInitOptions() RclInitOptions {
	var ret C.rcl_init_options_t = C.rcl_get_zero_initialized_init_options()
	return RclInitOptions(ret)
}

//
func RclGetDefaultAllocator() RclAllocator {
	var ret = C.rcl_get_default_allocator()
	return RclAllocator(ret)
}

//
func RclInitOptionsInit(opt *RclInitOptions, allocator RclAllocator) RclRet {
	var ret C.int32_t = C.rcl_init_options_init(
		(*C.rcl_init_options_t)(opt),
		(C.rcl_allocator_t)(allocator),
	)
	return RclRet(ret)
}

//
func RclInitOptionsFini(opt *RclInitOptions) RclRet {
	var ret C.int32_t = C.rcl_init_options_fini(
		(*C.rcl_init_options_t)(opt),
	)
	return RclRet(ret)
}

//
func RclInit(argc int, argv []string, opt *RclInitOptions, ctx RclContextPtr) RclRet {

	var arg **C.char = nil
	if len(argv) != 0 {
		cArgv := make([]*C.char, len(argv))
		for i := range argv {
			cstr := C.CString(argv[i])
			defer C.free(unsafe.Pointer(cstr))
			cArgv[i] = cstr
		}
		arg = &cArgv[0]
	}

	var ret = C.rcl_init(
		C.int(argc),
		arg,
		(*C.rcl_init_options_t)(opt),
		(*C.rcl_context_t)(ctx),
	)

	return RclRet(ret)
}

// RclShutdown represents Signal global shutdown of rcl.
func RclShutdown(ctx RclContextPtr) RclRet {
	cCtx := (*C.rcl_context_t)(ctx)
	ret := C.rcl_shutdown(cCtx)
	C.free_context(cCtx)
	return RclRet(ret)
}

//
type RmwMessageInfo C.rmw_message_info_t

//
type RmwRequestID C.rmw_request_id_t
