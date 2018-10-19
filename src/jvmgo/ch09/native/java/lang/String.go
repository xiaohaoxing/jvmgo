package lang

import (
	"jvmgo/ch09/native"
	"jvmgo/ch09/rtda"
	"jvmgo/ch09/rtda/heap"
)

func init() {
	native.Register("java/lang/String", "intern", "()Ljava/lang/String;", intern)
}

func intern(frame *rtda.Frame) {
	this := frame.LocalVars().GetThis()
	interned := heap.InternString(this)
	frame.OperandStack().PushRef(interned)
}
