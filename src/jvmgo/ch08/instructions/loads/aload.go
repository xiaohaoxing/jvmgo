package loads

import "jvmgo/ch08/instructions/base"
import "jvmgo/ch08/rtda"

/*
从局部变量表（LocalVars）加载到栈顶（stack.push）
*/
//load ref from local variables
type ALOAD struct{ base.Index8Instruction }

func (self *ALOAD) Execute(frame *rtda.Frame) {
	_aload(frame, self.Index)
}

type ALOAD_0 struct{ base.NoOperandsInstruction }

func (self *ALOAD_0) Execute(frame *rtda.Frame) {
	_aload(frame, 0)
}

type ALOAD_1 struct{ base.NoOperandsInstruction }

func (self *ALOAD_1) Execute(frame *rtda.Frame) {
	_aload(frame, 1)
}

type ALOAD_2 struct{ base.NoOperandsInstruction }

func (self *ALOAD_2) Execute(frame *rtda.Frame) {
	_aload(frame, 2)
}

type ALOAD_3 struct{ base.NoOperandsInstruction }

func (self *ALOAD_3) Execute(frame *rtda.Frame) {
	_aload(frame, 3)
}

//公用代码
func _aload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetRef(index)
	frame.OperandStack().PushRef(val)
}
