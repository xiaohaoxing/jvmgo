package constants

import "jvmgo/ch11/instructions/base"
import "jvmgo/ch11/rtda"

type NOP struct {
	base.NoOperandsInstruction
}

func (self *NOP) Execute(frame *rtda.Frame) {
	//do nothing
}
