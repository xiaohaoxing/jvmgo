package references

import (
	"jvmgo/ch07/instructions/base"
	"jvmgo/ch07/rtda"
	"jvmgo/ch07/rtda/heap"
)

/**
后续是4个字节不是2个字节，
2-常量池索引
1-参数需要的 slot 数量
1-值必须是0，向后兼容
*/
type INVOKE_INTERFACE struct {
	index uint
	// count uint8
	// zero uint8
}

func (self *INVOKE_INTERFACE) FetchOperands(reader *base.BytecodeReader) {
	self.index = uint(reader.ReadInt16())
	reader.ReadInt8() // the count
	reader.ReadInt8() // the zero
}

func (self *INVOKE_INTERFACE) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	methodRef := cp.GetConstant(self.index).(*heap.InterfaceMethodRef)
	resolvedMethod := methodRef.ResolvedInterfaceMethod()
	// static 或者 private 则异常
	if resolvedMethod.IsStatic() || resolvedMethod.IsPrivate() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	// 弹出 this引用，如果是 null 则抛出 NPE，如果该对象的类没有实现该接口则抛出异常。
	ref := frame.OperandStack().GetRefFromTop(resolvedMethod.ArgSlotCount() - 1)
	if ref == nil {
		panic("java.lang.NullPointerException")
	}
	if !ref.Class().IsImplements(methodRef.ResolvedClass()) {
		panic("java.lang.IncompatibleClassChangeError")
	}

	methodToBeInvoked := heap.LookupMethodInClass(ref.Class(), methodRef.Name(), methodRef.Descriptor())
	// 没找到函数或者找到的函数是 abstract 的
	if methodToBeInvoked == nil || methodToBeInvoked.IsAbstract() {
		panic("java.lang.AbstractMethodError")
	}
	// 非 public 的方法调用也不行
	if !methodToBeInvoked.IsPublic() {
		panic("java.lang.IllegalAccessError")
	}

	base.InvokeMethod(frame, methodToBeInvoked)
}
