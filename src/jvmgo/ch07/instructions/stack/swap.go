package stack

import "jvmgo/ch07/instructions/base"
import "jvmgo/ch07/rtda"

//交换栈顶2个变量的指令
type SWAP struct{ base.NoOperandsInstruction }

func (self *SWAP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
}
