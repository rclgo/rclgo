package cwrap

// #include <stdint.h>
import "C"

type RclRet C.int32_t

const (
	Ok                     = 0
	Error                  = 1
	Timeout                = 2
	BadAlloc               = 10
	RmwInvalidArgument     = 11
	AlreadyInit            = 100
	NotInit                = 101
	MismatchedRmwID        = 102
	TopicNameInvalid       = 103
	ServiceNameInvalid     = 104
	UnknownSubstitution    = 105
	NodeInvalid            = 200
	NodeInvalidName        = 201
	NodeInvalidNamespace   = 202
	PublisherInvalid       = 300
	SubscriptionInvalid    = 400
	SubscriptionTakeFailed = 401
	ClientInvalid          = 500
	ClientTakeFailed       = 501
	ServiceInvalid         = 600
	ServiceTakeFailed      = 601
	TimerInvalid           = 800
	TimerCanceled          = 801
	WaitSetInvalid         = 900
	WaitSetEmpty           = 901
	WaitSetFull            = 902
)

func (r RclRet) Error() string { return r.String() }

func (r RclRet) String() string {
	switch r {
	case 0:
		return "OK"
	case 1:
		return "Error"
	case 2:
		return "Timeout"
	case 10:
		return "Bad alloc"
	case 11:
		return "Invalid argument"

	case 100:
		return "Already init"
	case 101:
		return "Not init"
	case 102:
		return "Mismatched RMW ID"
	case 103:
		return "Topic name invalid"
	case 104:
		return "Service name invalid"
	case 105:
		return "Unknown substitution"
	case 200:
		return "Node invalid"
	case 201:
		return "Node invalid name"
	case 202:
		return "Node invalid namespace"
	case 300:
		return "Publisher invalid"
	case 400:
		return "Subscription invalid"
	case 401:
		return "Subscription take failed"
	case 500:
		return "Client invalid"
	case 501:
		return "Client take failed"
	case 600:
		return "Service invalid"
	case 601:
		return "Service take failed"
	case 800:
		return "Timer invalid"
	case 801:
		return "Timer canceled"
	case 900:
		return "Wait set invalid"
	case 901:
		return "Wait set empty"
	case 902:
		return "Wait set full"
	}

	return "Unknown"
}
