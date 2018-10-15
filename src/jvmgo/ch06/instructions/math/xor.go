package math

import "jvmgo/ch06/instructions/base"
import "jvmgo/ch06/rtda"

//int 异或运算
type IXOR struct{ base.NoOperandsInstruction }

func (self *IXOR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	var2 := stack.PopInt()
	var1 := stack.PopInt()
	result := var1 ^ var2
	stack.PushInt(result)
}

//long 异或运算
type LXOR struct{ base.NoOperandsInstruction }

func (self *LXOR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	var2 := stack.PopLong()
	var1 := stack.PopLong()
	result := var1 ^ var2
	stack.PushLong(result)
}
