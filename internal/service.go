package cwrap

// #cgo CFLAGS: -I/opt/ros/eloquent/include
// #include <rcl/rcl.h>
import "C"

import "unsafe"

type RclService C.rcl_service_t

// Options available for a rcl service.
type RclServiceOptions C.rcl_service_options_t

// RclGetZeroInitializedService returns a rcl_service_t struct with members set to `NULL`.
func RclGetZeroInitializedService() RclService {
	var ret C.rcl_service_t = C.rcl_get_zero_initialized_service()
	return RclService(ret)
}

// RclServiceInit initialize an rcl service.
func RclServiceInit(
	service *RclService,
	node *RclNode,
	typeSupport *RosidlServiceTypeSupport,
	serviceName string,
	options *RclServiceOptions,
) RclRet {
	cName := C.CString(serviceName)
	defer C.free(unsafe.Pointer(cName))
	var ret C.rcl_ret_t = C.rcl_service_init(
		(*C.rcl_service_t)(service),
		(*C.rcl_node_t)(node),
		(*C.rosidl_service_type_support_t)(typeSupport),
		cName,
		(*C.rcl_service_options_t)(options),
	)
	return RclRet(ret)
}

// RclServiceFini finalize a rcl_service_t.
func RclServiceFini(
	service *RclService,
	node *RclNode,
) RclRet {
	var ret C.rcl_ret_t = C.rcl_service_fini(
		(*C.rcl_service_t)(service),
		(*C.rcl_node_t)(node),
	)
	return RclRet(ret)
}

// RclServiceGetDefaultOptions returns the default service options in a rcl_service_options_t.
func RclServiceGetDefaultOptions() RclServiceOptions {
	var ret C.rcl_service_options_t = C.rcl_service_get_default_options()
	return RclServiceOptions(ret)
}

// RclTakeRequest take a pending ROS request using an rcl service.
func RclTakeRequest(
	service *RclService,
	requestHeader *RmwRequestID,
	rosRequest unsafe.Pointer,
) RclRet {
	var ret C.rcl_ret_t = C.rcl_take_request(
		(*C.rcl_service_t)(service),
		(*C.rmw_request_id_t)(requestHeader),
		rosRequest,
	)
	return RclRet(ret)
}

// RclSendResponse send a ROS response to a client using a service.
func RclSendResponse(
	service *RclService,
	responseHeader *RmwRequestID,
	rosResponse unsafe.Pointer,
) RclRet {
	var ret C.rcl_ret_t = C.rcl_send_response(
		(*C.rcl_service_t)(service),
		(*C.rmw_request_id_t)(responseHeader),
		rosResponse,
	)
	return RclRet(ret)
}

// RclServiceGetServiceName get the topic name for the service.
func RclServiceGetServiceName(
	service *RclService,
) string {
	var ret *C.char = C.rcl_service_get_service_name(
		(*C.rcl_service_t)(service),
	)
	return C.GoString(ret)
}

// RclServiceGetOptions returns the rcl service options.
func RclServiceGetOptions(
	service *RclService,
) *RclServiceOptions {
	var ret *C.rcl_service_options_t = C.rcl_service_get_options(
		(*C.rcl_service_t)(service),
	)
	return (*RclServiceOptions)(ret)
}

// RclServiceGetRmwHandle returns the rmw service handle.
func RclServiceGetRmwHandle(
	service *RclService,
) *RmwService {
	var ret *C.rmw_service_t = C.rcl_service_get_rmw_handle(
		(*C.rcl_service_t)(service),
	)
	return (*RmwService)(ret)
}

// RclServiceIsValid check that the service is valid.
func RclServiceIsValid(
	service *RclService,
) bool {
	var ret C.bool = C.rcl_service_is_valid(
		(*C.rcl_service_t)(service),
	)
	return bool(ret)
}
