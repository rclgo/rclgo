package rclgo

import (
	cwrap "github.com/rclgo/rclgo/internal"
)

type WaitSet struct {
	rclWaitSet *cwrap.RclWaitSet
}

//
func NewWaitSet() WaitSet {
	return newZeroInitializedWaitSet()
}

//
func newZeroInitializedWaitSet() WaitSet {
	zeroWaitset := cwrap.RclGetZeroInitializedWaitSet()
	return WaitSet{&zeroWaitset}
}

//
func (w *WaitSet) Init(
	numSubs, numGuards, numTimers, numClients, numServices, numEvents uint,
	ctx cwrap.RclContextPtr,
	allo cwrap.RclAllocator,
) error {
	ret := cwrap.RclWaitSetInit(
		w.rclWaitSet,
		numSubs,
		numGuards,
		numTimers,
		numClients,
		numServices,
		numEvents,
		ctx,
		allo,
	)
	if ret != cwrap.Ok {
		return NewErr("RclWaitSetInit", ret)
	}

	return nil
}

//
func (w *WaitSet) Fini() error {
	ret := cwrap.RclWaitSetFini(w.rclWaitSet)
	if ret != cwrap.Ok {
		return NewErr("RclWaitSetFini", ret)
	}

	return nil
}

//
func (w *WaitSet) GetAllocator(allocator *cwrap.RclAllocator) error {
	ret := cwrap.RclWaitSetGetAllocator(w.rclWaitSet, allocator)
	if ret != cwrap.Ok {
		return NewErr("RclWaitSetGetAllocator", ret)
	}

	return nil
}

//
func (w *WaitSet) AddSubscription(rclSubscription *cwrap.RclSubscription) error {
	ret := cwrap.RclWaitSetAddSubscription(w.rclWaitSet, rclSubscription, nil)
	if ret != cwrap.Ok {
		return NewErr("RclWaitSetAddSubscription", ret)
	}
	return nil
}

//
func (w *WaitSet) Clear() error {
	ret := cwrap.RclWaitSetClear(w.rclWaitSet)
	if ret != cwrap.Ok {
		return NewErr("RclWaitSetClear", ret)
	}
	return nil
}

//
func (w *WaitSet) Wait(timeout int64) error {
	var ret = cwrap.RclWait(w.rclWaitSet, timeout)
	if ret != cwrap.Ok {
		return NewErr("RclWaitSetClear", ret)
	}
	return nil
}
