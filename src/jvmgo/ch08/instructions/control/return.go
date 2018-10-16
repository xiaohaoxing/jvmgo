package control

import (
	"jvmgo/ch08/instructions/base"
	"jvmgo/ch08/rtda"
)

type RETURN struct{ base.NoOperandsInstruction }

func (self *RETURN) FetchOperands(reader *base.BytecodeReader) {}

func (self *RETURN) Execute(frame *rtda.Frame) {
	frame.Thread().PopFrame()
}

type ARETURN struct{ base.NoOperandsInstruction }

func (self *ARETURN) FetchOperands(reader *base.BytecodeReader) {}

func (self *ARETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokeFrame := thread.TopFrame()
	returnVal := currentFrame.OperandStack().PopRef()
	invokeFrame.OperandStack().PushRef(returnVal)
}

type DRETURN struct{ base.NoOperandsInstruction }

func (self *DRETURN) FetchOperands(reader *base.BytecodeReader) {}

func (self *DRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokeFrame := thread.TopFrame()
	returnVal := currentFrame.OperandStack().PopDouble()
	invokeFrame.OperandStack().PushDouble(returnVal)
}

type FRETURN struct{ base.NoOperandsInstruction }

func (self *FRETURN) FetchOperands(reader *base.BytecodeReader) {}

func (self *FRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokeFrame := thread.TopFrame()
	returnVal := currentFrame.OperandStack().PopFloat()
	invokeFrame.OperandStack().PushFloat(returnVal)
}

type IRETURN struct{ base.NoOperandsInstruction }

func (self *IRETURN) FetchOperands(reader *base.BytecodeReader) {}

func (self *IRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokeFrame := thread.TopFrame()
	returnVal := currentFrame.OperandStack().PopInt()
	invokeFrame.OperandStack().PushInt(returnVal)
}

type LRETURN struct{ base.NoOperandsInstruction }

func (self *LRETURN) FetchOperands(reader *base.BytecodeReader) {}

func (self *LRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokeFrame := thread.TopFrame()
	returnVal := currentFrame.OperandStack().PopLong()
	invokeFrame.OperandStack().PushLong(returnVal)
}
