package comparisons

import "jvmgo/ch09/instructions/base"
import "jvmgo/ch09/rtda"

//比较 float 指令1:FCMPG
type FCMPG struct{ base.NoOperandsInstruction }

func (self *FCMPG) Execute(frame *rtda.Frame) {
	_fcmp(frame, true)
}

type FCMPL struct{ base.NoOperandsInstruction }

func (self *FCMPL) Execute(frame *rtda.Frame) {
	_fcmp(frame, false)
}

//辅助方法，用一个 bool 标记用了哪个方法。
//如果不能比较，2种指令得到的结果不相同。如果能比较，2种指令得到的结果一定相同。
func _fcmp(frame *rtda.Frame, flag bool) {
	stack := frame.OperandStack()
	var2 := stack.PopFloat()
	var1 := stack.PopFloat()
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
