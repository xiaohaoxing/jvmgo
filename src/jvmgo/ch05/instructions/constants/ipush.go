package constants

import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"
/*
包括BIPUSH和SIPUSH2个获取byte/short型数据扩展成int并push栈顶
*/
type BIPUSH struct {val int8}

func (self *BIPUSH) FetchOperands(reader *base.BytecodeReader) {
	self.val = reader.ReadInt8()
}

func (self *BIPUSH) Execute(frame *rtda.Frame) {
	i := int32(self.val)
	frame.OperandStack().PushInt(i)
}

type SIPUSH struct {val int16}

func (self *SIPUSH) FetchOperands(reader *base.BytecodeReader) {
	self.val = reader.ReadInt16()
}

func (self *SIPUSH) Execute(frame *rtda.Frame) {
	i := int32(self.val)
	frame.OperandStack().PushInt(i)
}