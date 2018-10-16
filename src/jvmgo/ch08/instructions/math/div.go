package math

import "jvmgo/ch08/instructions/base"
import "jvmgo/ch08/rtda"

//double 类型的除法
type DDIV struct{ base.NoOperandsInstruction }

func (self *DDIV) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	var2 := stack.PopDouble()
	var1 := stack.PopDouble()
	result := var1 / var2
	stack.PushDouble(result)
}

//float 类型的除法
type FDIV struct{ base.NoOperandsInstruction }

func (self *FDIV) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	var2 := stack.PopFloat()
	var1 := stack.PopFloat()
	result := var1 / var2
	stack.PushFloat(result)
}

//int 类型的除法
type IDIV struct{ base.NoOperandsInstruction }

func (self *IDIV) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	var2 := stack.PopInt()
	var1 := stack.PopInt()
	if var2 == 0 {
		panic("java.lang.ArithmeticException: / by zero!")
	}
	result := var1 / var2
	stack.PushInt(result)
}

//long 类型的除法
type LDIV struct{ base.NoOperandsInstruction }

func (self *LDIV) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	var2 := stack.PopLong()
	var1 := stack.PopLong()
	if var2 == 0 {
		panic("java.lang.ArithmeticException: / by zero!")
	}
	result := var1 / var2
	stack.PushLong(result)
}
