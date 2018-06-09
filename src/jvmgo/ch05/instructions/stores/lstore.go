package stores

import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"

//store long into local variable
type LSTORE struct {base.Index8Instruction}
type LSTORE_0 struct {base.NoOperandsInstruction}
type LSTORE_1 struct {base.NoOperandsInstruction}
type LSTORE_2 struct {base.NoOperandsInstruction}
type LSTORE_3 struct {base.NoOperandsInstruction}

//公用函数
func _lstore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopLong()
	frame.LocalVars().SetLong(index, val)
}

func (self *LSTORE) Execute(frame *rtda.Frame) {
	_lstore(frame, uint(self.Index))
}

func (self *LSTORE_0) Execute(frame *rtda.Frame) {
	_lstore(frame, 0)
}

func (self *LSTORE_1) Execute(frame *rtda.Frame) {
	_lstore(frame, 1)
}

func (self *LSTORE_2) Execute(frame *rtda.Frame) {
	_lstore(frame, 2)
}

func (self *LSTORE_3) Execute(frame *rtda.Frame) {
	_lstore(frame, 3)
}