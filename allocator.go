package rclgo

import (
	cwrap "github.com/rclgo/rclgo/internal"
)

type Allocator struct {
	rclAllocator *cwrap.RclAllocator
}

func newZeroInitializedAllocator() Allocator {
	zeroAllocator := cwrap.RcutilsGetZeroInitializedAllocator()
	return Allocator{&zeroAllocator}
}

func newDefaultAllocator() Allocator {
	defAllocator := cwrap.RcutilsGetDefaultAllocator()
	return Allocator{&defAllocator}
}
