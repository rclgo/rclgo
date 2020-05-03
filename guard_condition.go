package rclgo

import (
	cwrap "github.com/rclgo/rclgo/internal"
)

type GuardCondition struct {
	context           Context
	rclGuardCondition *cwrap.RclGuardCondition
}

//
func NewGuardCondition() (GuardCondition, error) {
	ctx := GetDefaultContext()
	var gc = cwrap.RclGetZeroInitializedGuardCondition()
	var opt = cwrap.RclGuardConditionGetDefaultOptions()
	ret := cwrap.RclGuardConditionInit(&gc, ctx.rclContext, opt)
	if ret != cwrap.Ok {
		return GuardCondition{}, NewErr("cwrap.RclTriggerGuardCondtion", ret)
	}
	return GuardCondition{
		context:           ctx,
		rclGuardCondition: &gc,
	}, nil
}

//
func (gc *GuardCondition) Trigger() error {
	ret := cwrap.RclTriggerGuardCondition(gc.rclGuardCondition)
	if ret != cwrap.Ok {
		return NewErr("cwrap.RclTriggerGuardCondtion", ret)
	}
	return nil
}
