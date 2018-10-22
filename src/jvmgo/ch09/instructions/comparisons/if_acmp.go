package comparisons

import "jvmgo/ch09/instructions/base"
import "jvmgo/ch09/rtda"

type IF_ACMPEQ struct{ base.BranchInstruction }

func (self *IF_ACMPEQ) Execute(frame *rtda.Frame) {
	if _acmp(frame) {
		base.Branch(frame, self.Offset)
	}
}

type IF_ACMPNE struct{ base.BranchInstruction }

func (self *IF_ACMPNE) Execute(frame *rtda.Frame) {
	if !_acmp(frame) {
		base.Branch(frame, self.Offset)
	}
}


func _acmp(frame *rtda.Frame) bool {
	stack := frame.OperandStack()
	var1 := stack.PopRef()
	var2 := stack.PopRef()
	return var1 == var2
}