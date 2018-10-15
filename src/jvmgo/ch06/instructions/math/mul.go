package math

import "jvmgo/ch06/instructions/base"
import "jvmgo/ch06/rtda"

//double 类型的乘法
type DMUL struct{ base.NoOperandsInstruction }

func (self *DMUL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	var2 := stack.PopDouble()
	var1 := stack.PopDouble()
	result := var1 * var2
	stack.PushDouble(result)
}

//float 类型的乘法
type FMUL struct{ base.NoOperandsInstruction }

func (self *FMUL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	var2 := stack.PopFloat()
	var1 := stack.PopFloat()
	result := var1 * var2
	stack.PushFloat(result)
}

//int 类型的乘法
type IMUL struct{ base.NoOperandsInstruction }

func (self *IMUL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	var2 := stack.PopInt()
	var1 := stack.PopInt()
	result := var1 * var2
	stack.PushInt(result)
}

//long 类型的乘法
type LMUL struct{ base.NoOperandsInstruction }

func (self *LMUL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	var2 := stack.PopLong()
	var1 := stack.PopLong()
	result := var1 * var2
	stack.PushLong(result)
}
