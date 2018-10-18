package references

import (
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtda"
	"jvmgo/ch06/rtda/heap"
)

//NEW 指令
type NEW struct{ base.Index16Instruction }

func (self *NEW) Execute(frame *rtda.Frame) {
	//取出常量池
	cp := frame.Method().Class().ConstantPool()
	//得到该常量引用，转型为引用类型
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef)
	//解析成 class 对象
	class := classRef.ResolvedClass()
	//抽象类和接口不能实例化
	if class.IsInterface() || class.IsAbstract() {
		panic("java.lang.InstantiationError!")
	}
	ref := class.NewObject()
	//作为操作数推入栈中
	frame.OperandStack().PushRef(ref)
}
