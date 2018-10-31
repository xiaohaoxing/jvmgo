package lang

import (
	"jvmgo/ch10/native"
	"jvmgo/ch10/rtda"
	"jvmgo/ch10/rtda/heap"
)

func init() {
	native.Register("java/lang/System", "arraycopy", "(Ljava/lang/Object;ILjava/lang/Object;II)V", arraycopy)
}

// 实现 public static native void arraycopy(Object src, int srcPos, Object dest, int destPos, int length)
func arraycopy(frame *rtda.Frame) {
	// 获取所有参数
	vars := frame.LocalVars()
	src := vars.GetRef(0)
	srcPos := vars.GetInt(1)
	dest := vars.GetRef(2)
	destPos := vars.GetInt(3)
	length := vars.GetInt(4)
	// 源和目标数组都不能为 null
	if src == nil || dest == nil {
		panic("java.lang.NullPointerException")
	}
	// 源和目标数组必须兼容
	if !checkArrayCopy(src, dest) {
		panic("java.lang.ArrayStoreException")
	}
	// 检查位置异常
	if srcPos < 0 || destPos < 0 || length < 0 || srcPos+length > src.ArrayLength() || destPos+length > dest.ArrayLength() {
		panic("java.lang.IndexOutOfBoundsException")
	}

	// 进行拷贝
	heap.ArrayCopy(src, dest, srcPos, destPos, length)
}

func checkArrayCopy(src, dest *heap.Object) bool {
	srcClass := src.Class()
	destClass := dest.Class()
	if !srcClass.IsArray() || !destClass.IsArray() {
		return false
	}
	if srcClass.ComponentClass().IsPrimitive() || destClass.ComponentClass().IsPrimitive() {
		return srcClass == destClass
	}
	return true
}
