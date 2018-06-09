package stack

import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"

type DUP struct { base.NoOperandsInstruction }
type DUP_X1 struct { base.NoOperandsInstruction }
type DUP_X2 struct { base.NoOperandsInstruction }
type DUP2 struct { base.NoOperandsInstruction }
type DUP2_X1 struct { base.NoOperandsInstruction }
type DUP2_X2 struct { base.NoOperandsInstruction }

//栈顶：c->b->a
//		
//		c->b->a->a
func (self *DUP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot := stack.PopSlot()
	stack.PushSlot(slot)
	stack.PushSlot(slot)
}

//栈顶: c->b->a
//
//		c->a->b->a
func (self *DUP_X1) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

//栈顶：c->b->a
//
//		a->c->b->a
func (self *DUP_X2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	slot3 := stack.PopSlot()
	
}
//c->b->a
//
//c->b->a->b->a
func (self *DUP2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}
//c->b->a
//
//b->a->c->b->a
func (self *DUP2_X1) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	slot3 := stack.PopSlot()
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
	stack.PushSlot(slot3)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

//d->c->b->a
//
//b->a->d->c->b->a
func (self *DUP2_X2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	a := stack.PopSlot()
	b := stack.PopSlot()
	c := stack.PopSlot()
	d := stack.PopSlot()
	stack.PushSlot(b)
	stack.PushSlot(a)
	stack.PushSlot(d)
	stack.PushSlot(c)
	stack.PushSlot(b)
	stack.PushSlot(a)
}