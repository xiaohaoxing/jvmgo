package misc

import (
	"jvmgo/ch11/instructions/base"
	"jvmgo/ch11/native"
	"jvmgo/ch11/rtda"
)

func init() {
	native.Register("sun/misc/VM", "initialize", "()V", initialize)
}

/*等价于 Java 代码 */
// private static native void initialize() {
// 		VM.savedProps.setProperty("foo", "bar");
// 	}
func initialize(frame *rtda.Frame) {
	classLoader := frame.Method().Class().Loader()
	jlSysClass := classLoader.LoadClass("java/lang/System")
	initSysClass := jlSysClass.GetStaticMethod("initializeSystemClass", "()V")
	base.InvokeMethod(frame, initSysClass)
}
