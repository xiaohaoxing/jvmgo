package base

import (
	"jvmgo/ch07/rtda"
	"jvmgo/ch07/rtda/heap"
)

func InvokeMethod(invokerFrame *rtda.Frame, method *heap.Method) {
	thread := invokerFrame.Thread()
	newFrame := thread.NewFrame(method)
	thread.PushFrame(newFrame)

	argSlot := int(method.ArgSlotCount())
	if argSlot > 0 {
	}
	for i := argSlot - 1; i >= 0; i-- {
		slot := invokerFrame.OperandStack().PopSlot()
		newFrame.LocalVars().SetSlot(uint(i), slot)
	}
}