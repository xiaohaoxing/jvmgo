package lang

import (
	"jvmgo/ch09/native"
	"jvmgo/ch09/rtda"
)

func init() {
	native.Register("java/lang/Object", "getClass", "()Ljava/lang/Class;", getClass)
}

// 实现 public final native Class<?> getClass()
func getClass(frame *rtda.Frame) {
	this := frame.LocalVars().GetThis()
	class := this.Class().GetJClass()
	frame.OperandStack().PushRef(class)
}
