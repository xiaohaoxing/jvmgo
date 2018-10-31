package references

import (
	"jvmgo/ch10/instructions/base"
	"jvmgo/ch10/rtda"
	"jvmgo/ch10/rtda/heap"
)

//PUT_STATIC 指令
type PUT_STATIC struct{ base.Index16Instruction }

func (self *PUT_STATIC) Execute(frame *rtda.Frame) {
	//当前方法
	currentMethod := frame.Method()
	//当前类
	currentClass := currentMethod.Class()
	//当前类常量池
	cp := currentClass.ConstantPool()
	//静态成员的引用
	fieldRef := cp.GetConstant(self.Index).(*heap.FieldRef)
	//从引用得到静态成员
	field := fieldRef.ResolvedField()
	//得到该静态成员的类，可能这个类还没初始化（当前类的父类或者接口等）
	class := field.Class()

	if !class.InitStarted() {
		frame.RevertNextPC()
		base.InitClass(frame.Thread(), class)
		return
	}
	//通过类访问对象成员：运行时错误！
	//如：
	// class Apple{int weight;}
	// Apple.weight = 10
	//就会抛出异常。
	if !field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	//final 修饰的对象成员变量只能在该类的构造函数中赋值，否则抛出异常！
	//"<clinit>"：编译器生成的初始化方法的名字
	if field.IsFinal() {
		if currentClass != class || currentMethod.Name() != "<clinit>" {
			panic("java.lang.IllegalAccessError")
		}
	}
	//根据描述符得到类型，从栈中弹出操作数并赋值给局部变量表对应的变量
	descriptor := field.Descriptor()
	slotId := field.SlotId()
	slots := class.StaticVars()
	stack := frame.OperandStack()
	switch descriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I':
		slots.SetInt(slotId, stack.PopInt())
	case 'F':
		slots.SetFloat(slotId, stack.PopFloat())
	case 'D':
		slots.SetDouble(slotId, stack.PopDouble())
	case 'J':
		slots.SetLong(slotId, stack.PopLong())
	case 'L', '[':
		slots.SetRef(slotId, stack.PopRef())
	}
}
