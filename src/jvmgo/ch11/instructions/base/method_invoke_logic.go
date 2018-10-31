package base

import (
	"jvmgo/ch11/rtda"
	"jvmgo/ch11/rtda/heap"
)

func InvokeMethod(invokerFrame *rtda.Frame, method *heap.Method) {
	thread := invokerFrame.Thread()
	newFrame := thread.NewFrame(method)
	thread.PushFrame(newFrame)

	argSlot := int(method.ArgSlotCount())
	if argSlot > 0 {
		for i := argSlot - 1; i >= 0; i-- {
			slot := invokerFrame.OperandStack().PopSlot()
			newFrame.LocalVars().SetSlot(uint(i), slot)
		}
	}

	// Native method hack
	// 只有registerNatives 方法处理，其他的直接 panic
	//if method.IsNative() {
	//	if method.Name() == "registerNatives" {
	//		thread.PopFrame()
	//	} else {
	//		panic(fmt.Sprintf("native method: %v.%v%v\n", method.Class().Name(), method.Name(), method.Descriptor()))
	//	}
	//}

}
