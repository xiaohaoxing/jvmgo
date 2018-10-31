package stack

import "jvmgo/ch11/instructions/base"
import "jvmgo/ch11/rtda"

//弹出一个操作数的指令
type POP struct{ base.NoOperandsInstruction }

func (self *POP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
}

//弹出2个操作数的指令（针对 long 和 double）
type POP2 struct{ base.NoOperandsInstruction }

func (self *POP2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
	stack.PopSlot()
}
