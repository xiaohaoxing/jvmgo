package control

import "jvmgo/ch11/instructions/base"
import "jvmgo/ch11/rtda"

type TABLE_SWITCH struct {
	defaultOffset int32
	low           int32
	high          int32
	jumpOffsets   []int32
}

//如果在 low 和 high 之间，就修改offset为操作数 offset，否则 offset 设置为default 值
func (self *TABLE_SWITCH) Execute(frame *rtda.Frame) {
	index := frame.OperandStack().PopInt()
	var offset int
	if index >= self.low && index <= self.high {
		offset = int(self.jumpOffsets[index-self.low])
	} else {
		offset = int(self.defaultOffset)
	}
	base.Branch(frame, offset)
}

func (self *TABLE_SWITCH) FetchOperands(reader *base.BytecodeReader) {
	reader.SkipPadding()
	self.defaultOffset = reader.ReadInt32()
	self.low = reader.ReadInt32()
	self.high = reader.ReadInt32()
	count := self.high - self.low + 1
	self.jumpOffsets = reader.ReadInt32s(count)
}