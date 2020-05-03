package cwrap

// #cgo CFLAGS: -I/opt/ros/eloquent/include
// #include <rcl/rcl.h>
// #include <stdlib.h>
// // This is needed to avoid `panic: runtime error: cgo argument has Go pointer
// // to Go pointer` when passing an `rcl_node_t*` to a C function. `rcl_node_t`
// // contains a pointer to an `rcl_context_t` which must therefore be a C
// // pointer when calling C functions from Go. Remember to free this memory
// // later.
// rcl_context_t* get_zero_context_ptr() {
// 	rcl_context_t* ptr = (rcl_context_t*)malloc(sizeof(rcl_context_t));
// 	*ptr = rcl_get_zero_initialized_context();
// }
import "C"

//
type RclContextPtr *C.rcl_context_t

//
func GetZeroInitializedContextPtr() RclContextPtr {
	var ctxPtr (*C.rcl_context_t) = C.get_zero_context_ptr()
	return RclContextPtr(ctxPtr)
}
