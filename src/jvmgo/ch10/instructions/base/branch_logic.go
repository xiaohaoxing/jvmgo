package base

import "jvmgo/ch10/rtda"

//设置下一个指令的位置
func Branch(frame *rtda.Frame, offset int) {
	pc := frame.Thread().PC()
	nextPC := pc + offset
	frame.SetNextPC(nextPC)
}
