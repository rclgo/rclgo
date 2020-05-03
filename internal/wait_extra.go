package cwrap

// #cgo CFLAGS: -I/opt/ros/eloquent/include
// #include <rcl/rcl.h>
// rcl_subscription_t * subscription_at(rcl_subscription_t **subs, size_t ix) {
// 	return subs[ix];
// }
import "C"

//
func (ws *RclWaitSet) GetSubscriptions() []*RclSubscription {
	var n C.ulong = (*C.rcl_wait_set_t)(ws).size_of_subscriptions
	var subs **C.rcl_subscription_t = (*C.rcl_wait_set_t)(ws).subscriptions
	var ret = make([]*RclSubscription, n)
	for i := C.ulong(0); i < n; i++ {
		ret[i] = (*RclSubscription)(C.subscription_at(subs, i))
	}
	return ret
}
