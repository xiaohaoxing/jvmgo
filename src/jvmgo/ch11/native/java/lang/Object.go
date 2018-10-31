package lang

import (
	"jvmgo/ch11/native"
	"jvmgo/ch11/rtda"
	"unsafe"
)

func init() {
	native.Register("java/lang/Object", "getClass", "()Ljava/lang/Class;", getClass)
	native.Register("java/lang/Object", "hashCode", "()I", hashCode)
	native.Register("java/lang/Object", "clone", "()Ljava/lang/Object;", clone)
}

// 实现 public final native Class<?> getClass()
func getClass(frame *rtda.Frame) {
	this := frame.LocalVars().GetThis()
	class := this.Class().GetJClass()
	frame.OperandStack().PushRef(class)
}

// 仅仅把 Object 转为 uintptr 类型，再转成 int32 推入栈
func hashCode(frame *rtda.Frame) {
	this := frame.LocalVars().GetThis()
	hash := int32(uintptr(unsafe.Pointer(this)))
	frame.OperandStack().PushInt(hash)
}

func clone(frame *rtda.Frame) {
	this := frame.LocalVars().GetThis()

	cloneable := this.Class().Loader().LoadClass("java/lang/Cloneable")
	if !this.Class().IsImplements(cloneable) {
		panic("java.lang.CloneableNotSupportedException")
	}

	frame.OperandStack().PushRef(this.Clone())
}
