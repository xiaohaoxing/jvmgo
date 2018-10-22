package base

import (
	"jvmgo/ch09/rtda"
	"jvmgo/ch09/rtda/heap"
)

func InitClass(thread *rtda.Thread, class *heap.Class) {
	class.StartInit()
	scheduleClinit(thread, class)
	initSuperClass(thread, class)
}

func scheduleClinit(thread *rtda.Thread, class *heap.Class) {
	clinit := class.GetClinitMethod()
	if clinit != nil && clinit.Class() == class{
		newFrame := thread.NewFrame(clinit)
		thread.PushFrame(newFrame)
	}
}

// 一直向上找初始化，每次都把初始化的帧放在子类的初始化的帧上面，保证执行顺序
func initSuperClass(thread *rtda.Thread, class *heap.Class) {
	if !class.IsInterface() {
		superClass := class.SuperClass()
		if superClass != nil && superClass.InitStarted() {
			InitClass(thread, superClass)
		}
	}
}
