package main

import (
	"fmt"
	"jvmgo/ch06/rtda/heap"
)
import "jvmgo/ch06/instructions"
import "jvmgo/ch06/instructions/base"
import "jvmgo/ch06/rtda"

//执行一个方法。按顺序读取attr，local 变量，操作数栈，方法的字节码。
//初始化各个成员
//循环执行字节码
func interpret(method *heap.Method) {
	thread := rtda.NewThread()
	frame := thread.NewFrame(method)
	thread.PushFrame(frame)

	defer catchErr(frame)
	loop(thread, method.Code())
}

func catchErr(frame *rtda.Frame) {
	if r := recover(); r != nil {
		fmt.Printf("LocalVars:%v\n", frame.LocalVars())
		fmt.Printf("OperandStack:%v\n", frame.OperandStack())
		panic(r)
	}
}

func loop(thread *rtda.Thread, bytecode []byte) {
	frame := thread.PopFrame()
	reader := &base.BytecodeReader{}
	for {
		//计算 PC
		pc := frame.NextPC()
		thread.SetPC(pc)
		reader.Reset(bytecode, pc)
		//读取指令
		opcode := reader.ReadUint8()
		inst := instructions.NewInstruction(opcode)
		//读取操作数:FetchOperands
		inst.FetchOperands(reader)
		//设置下一个指令
		frame.SetNextPC(reader.PC())
		//执行:execute
		fmt.Printf("pc:%2d inst:%T %v\n", pc, inst, inst)
		inst.Execute(frame)
	}
}
