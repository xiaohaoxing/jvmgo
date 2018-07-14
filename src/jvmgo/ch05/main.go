package main

import "fmt"
import "strings"
import "jvmgo/ch05/classfile"
import "jvmgo/ch05/classpath"

func main() {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Println("version0.0.5")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	// frame := rtda.NewFrame(100, 100)
	// testLocalVars(frame.LocalVars())
	// testOperandStack(frame.OperandStack())
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	className := strings.Replace(cmd.class, ".", "/", -1)
	cf := loadClass(className, cp)
	mainMethod := getMainMethod(cf)
	if mainMethod != nil {
		interpret(mainMethod)
	}else {
		fmt.Printf("Main method not found in class %s\n", cmd.class)
	}
}

func loadClass(className string, cp *classpath.Classpath) *classfile.ClassFile {
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		panic(err)
	}
	cf, err := classfile.Parse(classData)
	if err != nil {
		panic(err)
	}
	return cf
} 

func getMainMethod(cf *classfile.ClassFile) *classfile.MemberInfo {
	for _, m := range cf.Methods() {
		if m.Name() == "main" && m.Descriptor() == "([Ljava/lang/String;)V" {
			return m
		}
	}
	return nil
}
// func testLocalVars(vars rtda.LocalVars) {
// 	vars.SetInt(0, 100)
// 	vars.SetInt(1, -100)
// 	vars.SetLong(2, 2997924580)
// 	vars.SetLong(4, -2997924580)
// 	vars.SetFloat(6, 3.1415926)
// 	vars.SetDouble(7, 2.77556345)
// 	vars.SetRef(9, nil)

// 	println(vars.GetInt(0))
// 	println(vars.GetInt(1))
// 	println(vars.GetLong(2))
// 	println(vars.GetLong(4))
// 	println(vars.GetFloat(6))
// 	println(vars.GetDouble(7))
// 	println(vars.GetRef(9))
// }

// func testOperandStack(ops *rtda.OperandStack) {
// 	ops.PushInt(100)
// 	ops.PushInt(-100)
// 	ops.PushLong(2997924580)
// 	ops.PushLong(-2997924580)
// 	ops.PushFloat(3.1415926)
// 	ops.PushDouble(2.77556345)
// 	ops.PushRef(nil)

// 	println(ops.PopRef())
// 	println(ops.PopDouble())
// 	println(ops.PopFloat())
// 	println(ops.PopLong())
// 	println(ops.PopLong())
// 	println(ops.PopInt())
// 	println(ops.PopInt())
// }