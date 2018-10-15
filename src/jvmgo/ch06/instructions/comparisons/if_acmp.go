package comparisons

import "jvmgo/ch06/instructions/base"
import "jvmgo/ch06/rtda"

type IF_ACMPEQ struct{ base.BranchInstruction }

func (self *IF_ACMPEQ) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	var2 := stack.PopRef()
	var1 := stack.PopRef()
	if var1 == var2 {
		base.Branch(frame, self.Offset)
	}
}

type IF_ACMPNE struct{ base.BranchInstruction }

func (self *IF_ACMPNE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	var2 := stack.PopRef()
	var1 := stack.PopRef()
	if var1 != var2 {
		base.Branch(frame, self.Offset)
	}
}
