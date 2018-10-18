package stack

import "jvmgo/ch09/instructions/base"
import "jvmgo/ch09/rtda"

//交换栈顶2个变量的指令
type SWAP struct{ base.NoOperandsInstruction }

func (self *SWAP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
}
