package comparisons

import "jvmgo/ch10/instructions/base"
import "jvmgo/ch10/rtda"

//比较 double 指令1:DCMPG
type DCMPG struct{ base.NoOperandsInstruction }

func (self *DCMPG) Execute(frame *rtda.Frame) {
	_dcmp(frame, true)
}

//比较 double 指令2:DCMPL
type DCMPL struct{ base.NoOperandsInstruction }

func (self *DCMPL) Execute(frame *rtda.Frame) {
	_dcmp(frame, false)
}

//辅助方法，用一个 bool 标记用了哪个方法。
//如果不能比较，2种指令得到的结果不相同。如果能比较，2种指令得到的结果一定相同。
func _dcmp(frame *rtda.Frame, flag bool) {
	stack := frame.OperandStack()
	var2 := stack.PopDouble()
	var1 := stack.PopDouble()
	if var1 > var2 {
		stack.PushInt(1)
	} else if var1 == var2 {
		stack.PushInt(0)
	} else if var1 < var2 {
		stack.PushInt(-1)
	} else if flag {
		stack.PushInt(1)
	} else {
		stack.PushInt(-1)
	}
}
