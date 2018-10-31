package lang

import (
	"jvmgo/ch09/native"
	"jvmgo/ch09/rtda"
	"jvmgo/ch09/rtda/heap"
)

func init() {
	native.Register("java/lang/Class", "getPrimitiveClass", "(Ljava/lang/String;)Ljava/lang/Class;", getPrimitiveClass)
	native.Register("java/lang/Class", "getName0", "()Ljava/lang/String;", getName0)
	native.Register("java/lang/Class", "desiredAssertionStatus0", "(Ljava/lang/Class;)Z", desiredAssertionStatus0)
}
func getPrimitiveClass(frame *rtda.Frame) {
	nameObj := frame.LocalVars().GetRef(0)
	name := heap.GoString(nameObj)

	loader := frame.Method().Class().Loader()
	class := loader.LoadClass(name).GetJClass()

	frame.OperandStack().PushRef(class)
}

// 实现 private native String getName0()
func getName0(frame *rtda.Frame) {
	this := frame.LocalVars().GetThis()
	class := this.GetExtra().(*heap.Class)

	name := class.JavaName()
	nameObj := heap.JString(class.Loader(), name)

	frame.OperandStack().PushRef(nameObj)
}

// 实现 private static native boolean desiredAssertionStatus(Class<?> clazz)
func desiredAssertionStatus0(frame *rtda.Frame) {
	frame.OperandStack().PushBoolean(false)
}
