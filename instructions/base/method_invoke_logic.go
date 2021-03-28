package base

import (
	"jvmgo/rtda"
	"jvmgo/rtda/heap"
)

func InvokeMethod(invokerFrame *rtda.Frame, method *heap.Method) *rtda.Frame {
	thread := invokerFrame.Thread()
	newFrame := thread.NewFrame(method)
	thread.PushFrame(newFrame)

	argSlotCount := int(method.ArgSlotCount())
	if argSlotCount > 0 {
		for i := argSlotCount - 1; i >= 0; i-- {
			slot := invokerFrame.OperandStack().PopSlot()
			newFrame.LocalVars().SetSlot(uint(i), slot)
		}
	}
	return newFrame
}
