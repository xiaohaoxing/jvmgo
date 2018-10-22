package main

import (
	"fmt"
	"jvmgo/ch09/rtda/heap"
)
import "jvmgo/ch09/instructions"
import "jvmgo/ch09/instructions/base"
import "jvmgo/ch09/rtda"

//执行一个方法。按顺序读取attr，local 变量，操作数栈，方法的字节码。
//初始化各个成员
//循环执行字节码
func interpret(method *heap.Method, logInst bool, args []string) {
	thread := rtda.NewThread()
	frame := thread.NewFrame(method)
	thread.PushFrame(frame)
	// 把命令行参数转化成 Java 字符串数组
	jArgs := createArgsArray(method.Class().Loader(), args)
	frame.LocalVars().SetRef(0, jArgs)

	defer catchErr(thread)
	loop(thread, logInst)
}


func createArgsArray(loader *heap.ClassLoader, args []string) *heap.Object {
	stringClass := loader.LoadClass("java/lang/String")
	argsArr := stringClass.ArrayClass().NewArray(uint(len(args)))
	jArgs := argsArr.Refs()
	for i, arg := range args {
		jArgs[i] = heap.JString(loader, arg)
	}
	return argsArr
}


func catchErr(thread *rtda.Thread) {
	if r := recover(); r != nil {
		logFrames(thread)
		panic(r)
	}
}

func loop(thread *rtda.Thread, logInst bool) {
	//frame := thread.PopFrame()
	reader := &base.BytecodeReader{}
	for {
		frame := thread.CurrentFrame()
		//计算 PC
		pc := frame.NextPC()
		thread.SetPC(pc)
		reader.Reset(frame.Method().Code(), pc)
		//读取指令
		opcode := reader.ReadUint8()
		inst := instructions.NewInstruction(opcode)
		//读取操作数:FetchOperands
		inst.FetchOperands(reader)
		//设置下一个指令
		frame.SetNextPC(reader.PC())
		// 从命令行过来的，是否记录指令
		if logInst {
			logInstruction(frame, inst)
		}
		//执行:execute
		//fmt.Printf("pc:%2d inst:%T %v\n", pc, inst, inst)
		inst.Execute(frame)
		if thread.IsStackEmpty() {
			break
		}
	}
}

// 异常处理，把 thread 里的所有 frame 的信息打印出来
func logFrames(thread *rtda.Thread) {
	for !thread.IsStackEmpty() {
		frame := thread.PopFrame()
		method := frame.Method()
		className := method.Class().Name()
		fmt.Printf(">> pc:%4d %v.%v%v \n", frame.NextPC(), className, method.Name(), method.Descriptor())
	}
}

func logInstruction(frame *rtda.Frame, inst base.Instruction) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	pc := frame.Thread().PC()
	fmt.Printf("%v.%v() #%2d %T %v\n", className, methodName, pc, inst, inst)
}