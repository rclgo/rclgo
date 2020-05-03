package rclgo

import (
	cwrap "github.com/rclgo/rclgo/internal"
)

// Err implements the Error interface for rcl errors.
type Err struct {
	msg   string
	ret   cwrap.RclRet
	state *cwrap.RcutilsErrorState
}

// NewErr gets the error state from rcutils and wraps it in an Err.
// Intended for internal use only.
func NewErr(msg string, ret cwrap.RclRet) Err {
	return Err{
		msg:   msg,
		ret:   ret,
		state: cwrap.RcutilsGetErrorState(),
	}
}

func (e Err) Error() string {
	if e.ret == cwrap.Ok {
		return e.msg
	}
	return e.msg + ": " + e.ret.String() + "\n" + e.state.Error()
}

// Unwrap returns the underlying rcl error code.
func (e Err) Unwrap() error { return e.ret }
