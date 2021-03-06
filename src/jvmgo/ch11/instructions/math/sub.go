package math

import "jvmgo/ch11/instructions/base"
import "jvmgo/ch11/rtda"

//double 类型的减法
type DSUB struct{ base.NoOperandsInstruction }

func (self *DSUB) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	var2 := stack.PopDouble()
	var1 := stack.PopDouble()
	result := var1 - var2
	stack.PushDouble(result)
}

//float 类型的减法
type FSUB struct{ base.NoOperandsInstruction }

func (self *FSUB) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	var2 := stack.PopFloat()
	var1 := stack.PopFloat()
	result := var1 - var2
	stack.PushFloat(result)
}

//int 类型的减法
type ISUB struct{ base.NoOperandsInstruction }

func (self *ISUB) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	var2 := stack.PopInt()
	var1 := stack.PopInt()
	result := var1 - var2
	stack.PushInt(result)
}

//long 类型的减法
type LSUB struct{ base.NoOperandsInstruction }

func (self *LSUB) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	var2 := stack.PopLong()
	var1 := stack.PopLong()
	result := var1 - var2
	stack.PushLong(result)
}
