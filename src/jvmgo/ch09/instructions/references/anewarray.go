package references

import (
	"jvmgo/ch09/instructions/base"
	"jvmgo/ch09/rtda"
	"jvmgo/ch09/rtda/heap"
)

type ANEW_ARRAY struct {
	base.Index16Instruction
}

func (self *ANEW_ARRAY) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef)
	componentClass := classRef.ResolvedClass()
	// 如果数组的成员没有初始化，则中断数组的初始化，先初始化成员对象。
	// if componentClass.InitializationNotStarted() {
	// 	thread := frame.Thread()
	// 	frame.SetNextPC(thread.PC()) // undo anewarray
	// 	thread.InitClass(componentClass)
	// 	return
	// }


	stack := frame.OperandStack()
	count := stack.PopInt()
	if count < 0 {
		panic("java.lang.NegativeArraySizeException")
	}
	arrClass := componentClass.ArrayClass()
	arr := arrClass.NewArray(uint(count))
	stack.PushRef(arr)
}
