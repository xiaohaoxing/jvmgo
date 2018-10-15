package constants

import "jvmgo/ch07/instructions/base"
import "jvmgo/ch07/rtda"

/*
包括 BIPUSH 和 SIPUSH 2个获取 byte/short 型数据扩展成 int 并 push 栈顶
*/
type BIPUSH struct{ val int8 }

func (self *BIPUSH) FetchOperands(reader *base.BytecodeReader) {
	self.val = reader.ReadInt8()
}

func (self *BIPUSH) Execute(frame *rtda.Frame) {
	i := int32(self.val)
	frame.OperandStack().PushInt(i)
}

type SIPUSH struct{ val int16 }

func (self *SIPUSH) FetchOperands(reader *base.BytecodeReader) {
	self.val = reader.ReadInt16()
}

func (self *SIPUSH) Execute(frame *rtda.Frame) {
	i := int32(self.val)
	frame.OperandStack().PushInt(i)
}
