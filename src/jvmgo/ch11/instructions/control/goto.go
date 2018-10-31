package control

import "jvmgo/ch11/instructions/base"
import "jvmgo/ch11/rtda"

type GOTO struct{ base.BranchInstruction }

func (self *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.Offset)
}
