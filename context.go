package rclgo

import (
	cwrap "github.com/rclgo/rclgo/internal"
)

// Context encapsulates the non-global state of an init/shutdown cycle.
type Context struct {
	rclContext cwrap.RclContextPtr
	init       bool
}

var context Context

func GetDefaultContext() Context {
	if context.init == false {
		context = newZeroInitializedContext()
		context.Init()
		context.init = true
	}
	return context
}

// newZeroInitializedContext returns a zero initialization context object.
func newZeroInitializedContext() Context {
	ctxPtr := cwrap.GetZeroInitializedContextPtr()
	return Context{rclContext: ctxPtr, init: true}
}

// Init finalizes a context.
func (ctx *Context) Init() error {

	var opts = cwrap.RclGetZeroInitializedInitOptions()
	alloc := cwrap.RclGetDefaultAllocator()

	ret := cwrap.RclInitOptionsInit(&opts, alloc)
	if ret != cwrap.Ok {
		return NewErr("RclInitOptionsInit", ret)
	}

	ret = cwrap.RclInit(
		0,
		[]string{},
		&opts,
		ctx.rclContext,
	)
	if ret != cwrap.Ok {
		return NewErr("RclInit", ret)
	}

	ret = cwrap.RclInitOptionsFini(&opts)
	if ret != cwrap.Ok {
		return NewErr("RclInitOptionsFini", ret)
	}

	return nil
}

// Fini finalizes a context.
func (ctx *Context) Fini() error {
	ret := cwrap.RclContextFini(ctx.rclContext)
	if ret != cwrap.Ok {
		return NewErr("RclContextFini", ret)
	}

	return nil
}

// Shutdown shuts down the context.
func (ctx *Context) Shutdown() error {
	ret := cwrap.RclShutdown(ctx.rclContext)
	if ret != cwrap.Ok {
		return NewErr("RclShutdown", ret)
	}

	return nil
}

// IsValid returns `true` if the context is currently valid, otherwise `false`.
func (ctx *Context) IsValid() bool {
	ret := cwrap.RclContextIsValid(ctx.rclContext)
	return bool(ret)
}
