package loads

import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"
/*
从局部变量表（LocalVars）加载到栈顶（stack.push）
*/
//load int from local variables
type ILOAD struct {base.NoOperandsInstruction}
type ILOAD_0 struct {base.NoOperandsInstruction}
type ILOAD_1 struct {base.NoOperandsInstruction}
type ILOAD_2 struct {base.NoOperandsInstruction}
type ILOAD_3 struct {base.NoOperandsInstruction}

//公用代码
func _iload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetInt(index)
	frame.OperandStack().PushInt(val)
}

func (self *ILOAD) Execute(frame *rtda.Frame) {
	_iload(frame, uint(self.index))
}

func (self *ILOAD_0) Execute(frame *rtda.Frame) {
	_iload(frame, 0)
}

func (self *ILOAD_1) Execute(frame *rtda.Frame) {
	_iload(frame, 1)
}

func (self *ILOAD_2) Execute(frame *rtda.Frame) {
	_iload(frame, 2)
}

func (self *ILOAD_3) Execute(frame *rtda.Frame) {
	_iload(frame, 3)
}