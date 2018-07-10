package stack

import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"
//复制栈顶变量的操作系列
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
	a := stack.PopSlot()
	stack.PushSlot(a)
	stack.PushSlot(a)
}

//栈顶: c->b->a
//
//		c->b->a->b
func (self *DUP_X1) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	a := stack.PopSlot()
	b := stack.PopSlot()
	stack.PushSlot(b)
	stack.PushSlot(a)
	stack.PushSlot(b)
}

//栈顶：c->b->a
//
//		c->b->a->c
func (self *DUP_X2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	a := stack.PopSlot()
	b := stack.PopSlot()
	c := stack.PopSlot()
	stack.PushSlot(c)
	stack.PushSlot(b)
	stack.PushSlot(a)
	stack.PushSlot(c)
}
//c->b->a
//复制2个变量，无间隔
//c->b->a->b->a
func (self *DUP2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	a := stack.PopSlot()
	b := stack.PopSlot()
	stack.PushSlot(b)
	stack.PushSlot(a)
	stack.PushSlot(b)
	stack.PushSlot(a)
}
//c->b->a
//
//c->b->a->c->b
func (self *DUP2_X1) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	a := stack.PopSlot()
	b := stack.PopSlot()
	c := stack.PopSlot()
	stack.PushSlot(c)
	stack.PushSlot(b)
	stack.PushSlot(a)
	stack.PushSlot(c)
	stack.PushSlot(b)
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