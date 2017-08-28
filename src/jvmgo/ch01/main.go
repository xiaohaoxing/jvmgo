package main

import "fmt"
//启动时根据情况，有version就输出version，有help或class为空则输出用法，其他的情况则尝试启动jvm
func main() {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Println("version0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}
//假装JVM已经启动了。
func startJVM(cmd *Cmd) {
	fmt.Printf("classpath:%s class:%s args:%v\n", cmd.cpOption, cmd.class, cmd.args)
}