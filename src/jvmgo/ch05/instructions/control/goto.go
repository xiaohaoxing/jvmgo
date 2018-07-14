package control

import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"

type GOTO struct { base.BranchInstruction }

func (self *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.Offset)
}