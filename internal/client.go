package cwrap

// #cgo CFLAGS: -I/opt/ros/eloquent/include
// #include <rcl/rcl.h>
import "C"

import "unsafe"

type RclClient C.rcl_client_t

// Options available for a rcl_client_t.
type RclClientOptions C.rcl_client_options_t

// RclGetZeroInitializedClient returns a rcl_client_t struct with members set to `NULL`.
func RclGetZeroInitializedClient() RclClient {
	var ret C.rcl_client_t = C.rcl_get_zero_initialized_client()
	return RclClient(ret)
}

// RclClientInit initialize an rcl client.
func RclClientInit(
	client *RclClient,
	node *RclNode,
	typeSupport *RosidlServiceTypeSupport,
	serviceName string,
	options *RclClientOptions,
) RclRet {
	cName := C.CString(serviceName)
	defer C.free(unsafe.Pointer(cName))
	var ret C.rcl_ret_t = C.rcl_client_init(
		(*C.rcl_client_t)(client),
		(*C.rcl_node_t)(node),
		(*C.rosidl_service_type_support_t)(typeSupport),
		(*C.char)(cName),
		(*C.rcl_client_options_t)(options),
	)
	return RclRet(ret)
}

// RclClientFini finalize a rcl_client_t.
func RclClientFini(
	client *RclClient,
	node *RclNode,
) RclRet {
	var ret C.rcl_ret_t = C.rcl_client_fini(
		(*C.rcl_client_t)(client),
		(*C.rcl_node_t)(node),
	)
	return RclRet(ret)
}

// RclClientGetDefaultOptions returns the default client options in a rcl_client_options_t.
func RclClientGetDefaultOptions() RclClientOptions {
	var ret C.rcl_client_options_t = C.rcl_client_get_default_options()
	return RclClientOptions(ret)
}

// RclSendRequest send a ROS request using a client.
func RclSendRequest(
	client *RclClient,
	rosRequest unsafe.Pointer,
	sequenceNumber *int64,
) RclRet {
	var ret C.rcl_ret_t = C.rcl_send_request(
		(*C.rcl_client_t)(client),
		rosRequest,
		(*C.int64_t)(sequenceNumber),
	)
	return RclRet(ret)
}

// RclTakeResponse take a ROS response using a client
func RclTakeResponse(
	client *RclClient,
	requestHeader *RmwRequestID,
	rosResponse unsafe.Pointer,
) RclRet {
	var ret C.rcl_ret_t = C.rcl_take_response(
		(*C.rcl_client_t)(client),
		(*C.rmw_request_id_t)(requestHeader),
		rosResponse,
	)
	return RclRet(ret)
}

// RclClientGetServiceName get the name of the service that this client will request a response from.
func RclClientGetServiceName(
	client *RclClient,
) string {
	var ret *C.char = C.rcl_client_get_service_name(
		(*C.rcl_client_t)(client),
	)
	return C.GoString(ret)
}

// RclClientGetOptions returns the rcl client options.
func RclClientGetOptions(
	client *RclClient,
) *RclClientOptions {
	var ret *C.rcl_client_options_t = C.rcl_client_get_options(
		(*C.rcl_client_t)(client),
	)
	return (*RclClientOptions)(ret)
}

// // RclClientGetRmwHandle returns the rmw client handle.
// func RclClientGetRmwHandle(
// 	client *RclClient,
// ) RmwClient {
// 	var ret *C.rmw_client_t = C.rcl_client_get_rmw_handle(
// 		(*C.rcl_client_t)(client),
// 	)
// 	return RmwClient(ret)
// }

// RclClientIsValid check that the client is valid.
func RclClientIsValid(
	client *RclClient,
) bool {
	var ret C.bool = C.rcl_client_is_valid(
		(*C.rcl_client_t)(client),
	)
	return bool(ret)
}
