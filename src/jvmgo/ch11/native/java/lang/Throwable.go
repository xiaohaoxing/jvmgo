package lang

import (
	"jvmgo/ch11/native"
	"jvmgo/ch11/rtda"
	"jvmgo/ch11/rtda/heap"
)

type StackTraceElement struct {
	fileName   string // 所在文件名
	className  string // 声明方法的类名
	methodName string // 方法名
	lineNumber int    // 正在执行哪行代码
}

func init() {
	native.Register("java/lang/Throwable", "fillInStackTrace", "(I)Ljava/lang/Throwable;", fillInStackTrace)
}

func fillInStackTrace(frame *rtda.Frame) {
	this := frame.LocalVars().GetThis()
	frame.OperandStack().PushRef(this)

	stes := createStackTraceElements(this, frame.Thread())
	this.SetExtra(stes)
}

func createStackTraceElements(tObj *heap.Object, thread *rtda.Thread) []*StackTraceElement {
	// 需要跳过 fillInStackTrace(int) 和 fillInStackTrace() 2个方法，以及异常类的构造函数的所有层次。见下一个函数
	skip := distanceToObject(tObj.Class()) + 2
	frames := thread.GetFrames()[skip:]
	stes := make([]*StackTraceElement, len(frames))
	for i, frame := range frames {
		stes[i] = createStackTraceElement(frame)
	}
	return stes
}

// 异常类的构造函数，具体需要跳过多少个帧则根据这个类的继承层次决定的。
func distanceToObject(class *heap.Class) int {
	distance := 0

	for c := class.SuperClass(); c != nil; c = c.SuperClass() {
		distance++
	}
	return distance
}

// 创建一个异常堆栈信息
func createStackTraceElement(frame *rtda.Frame) *StackTraceElement {
	method := frame.Method()
	class := method.Class()
	return &StackTraceElement{
		fileName:   class.SourceFile(),
		className:  class.JavaName(),
		methodName: method.Name(),
		lineNumber: method.GetLineNumber(frame.NextPC() - 1),
	}
}
