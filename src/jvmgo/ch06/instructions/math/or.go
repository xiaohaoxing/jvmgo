package math

import "jvmgo/ch06/instructions/base"
import "jvmgo/ch06/rtda"

//int 类型的或运算
type IOR struct{ base.NoOperandsInstruction }

func (self *IOR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	var2 := stack.PopInt()
	var1 := stack.PopInt()
	result := var1 | var2
	stack.PushInt(result)
}

//long 类型的或运算
type LOR struct{ base.NoOperandsInstruction }

func (self *LOR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	var2 := stack.PopLong()
	var1 := stack.PopLong()
	result := var1 | var2
	stack.PushLong(result)
}
