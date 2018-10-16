package math

import "jvmgo/ch08/instructions/base"
import "jvmgo/ch08/rtda"

//double 类型的 add
type DADD struct{ base.NoOperandsInstruction }

func (self *DADD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	var1 := stack.PopDouble()
	var2 := stack.PopDouble()
	result := var1 + var2
	stack.PushDouble(result)
}

//float 类型的 add
type FADD struct{ base.NoOperandsInstruction }

func (self *FADD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	var1 := stack.PopFloat()
	var2 := stack.PopFloat()
	result := var1 + var2
	stack.PushFloat(result)
}

//int 类型的 add
type IADD struct{ base.NoOperandsInstruction }

func (self *IADD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	var1 := stack.PopInt()
	var2 := stack.PopInt()
	result := var1 + var2
	stack.PushInt(result)
}

//long 类型的 add
type LADD struct{ base.NoOperandsInstruction }

func (self *LADD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	var1 := stack.PopLong()
	var2 := stack.PopLong()
	result := var1 + var2
	stack.PushLong(result)
}
