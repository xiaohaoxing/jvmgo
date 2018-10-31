package comparisons

import "jvmgo/ch10/instructions/base"
import "jvmgo/ch10/rtda"

type IF_ICMPEQ struct{ base.BranchInstruction }

func (self *IF_ICMPEQ) Execute(frame *rtda.Frame) {
	var1, var2 := _icmpPop(frame)
	if var2 == var1 {
		base.Branch(frame, self.Offset)
	}
}

type IF_ICMPNE struct{ base.BranchInstruction }

func (self *IF_ICMPNE) Execute(frame *rtda.Frame) {
	var1, var2 := _icmpPop(frame)
	if var2 != var1 {
		base.Branch(frame, self.Offset)
	}
}

type IF_ICMPLT struct{ base.BranchInstruction }

func (self *IF_ICMPLT) Execute(frame *rtda.Frame) {
	var1, var2 := _icmpPop(frame)
	if var1 < var2 {
		base.Branch(frame, self.Offset)
	}
}

type IF_ICMPLE struct{ base.BranchInstruction }

func (self *IF_ICMPLE) Execute(frame *rtda.Frame) {
	var1, var2 := _icmpPop(frame)
	if var1 <= var2 {
		base.Branch(frame, self.Offset)
	}
}

type IF_ICMPGT struct{ base.BranchInstruction }

func (self *IF_ICMPGT) Execute(frame *rtda.Frame) {
	var1, var2 := _icmpPop(frame)
	if var1 > var2 {
		base.Branch(frame, self.Offset)
	}
}

type IF_ICMPGE struct{ base.BranchInstruction }

func (self *IF_ICMPGE) Execute(frame *rtda.Frame) {
	var1, var2 := _icmpPop(frame)
	if var1 >= var2 {
		base.Branch(frame, self.Offset)
	}
}

//共同方法，读取2个 int
func _icmpPop(frame *rtda.Frame) (var1, var2 int32) {
	stack := frame.OperandStack()
	var2 = stack.PopInt()
	var1 = stack.PopInt()
	return
}
