package comparisons

import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"

type IFEQ struct { base.NoOperandsInstruction }

func (self *INEQ) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val == 0 {
		base.Branch(frame, self.Offset)
	}
}

type IFNE struct { base.NoOperandsInstruction }

func (self *IFNE) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val != 0{
		base.Branch(frame, self.Offset)
	}
}

type IFLT struct { base.NoOperandsInstruction }

func (self *IFLT) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val < 0 {
		base.Branch(frame, self.Offset)
	}
}

type IFLE struct { base.NoOperandsInstruction }

func (self *IFLE) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val <= 0 {
		base.Branch(frame, self.Offset)
	}
}

type IFGT struct { base.NoOperandsInstruction }

func (self *IFGT) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val > 0 {
		base.Branch(frame, self.Offset)
	}
}

type IFGE struct { base.NoOperandsInstruction }
 
func (self *IFGE) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val >= 0 {
		base.Branch(frame, self.Offset)
	}
}