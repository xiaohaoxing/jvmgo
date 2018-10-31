package main

import (
	"fmt"
)

func main() {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Println("version0.0.6")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		newJVM(cmd).start()
	}
}

//func startJVM(cmd *Cmd) {
//	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
//	classLoader := heap.NewClassLoader(cp, cmd.verboseClassFlag)
//
//	className := strings.Replace(cmd.class, ".", "/", -1)
//	mainClass := classLoader.LoadClass(className)
//	mainMethod := mainClass.GetMainMethod()
//	if mainMethod != nil {
//		interpret(mainMethod, cmd.verboseInstFlag, cmd.args)
//	} else {
//		fmt.Printf("Main method not found in class %s\n", cmd.class)
//	}
//}
