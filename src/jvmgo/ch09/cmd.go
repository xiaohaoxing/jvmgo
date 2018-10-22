package main

import "flag"
import "fmt"
import "os"

//加上Xjre命令行参数
type Cmd struct {
	helpFlag    		bool
	versionFlag 		bool
	// 是否把类加载信息输出到 cmd
	verboseClassFlag 	bool
	// 是否把指令加载信息输出到 cmd
	verboseInstFlag 	bool
	cpOption        	string
	XjreOption      	string
	class           	string
	args            	[]string
}

func parseCmd() *Cmd {
	cmd := &Cmd{}

	flag.Usage = printUsage
	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.helpFlag, "?", false, "print help message")
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version and exit")
	flag.BoolVar(&cmd.verboseClassFlag, "verbose", false, "print the class load info")
	flag.BoolVar(&cmd.verboseClassFlag, "verbose:class", false, "print the class load info")
	flag.BoolVar(&cmd.verboseInstFlag, "verbose:inst", false, "print the instruction execute info")
	flag.StringVar(&cmd.cpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath")
	flag.StringVar(&cmd.XjreOption, "Xjre", "", "path to jre")
	flag.Parse()

	args := flag.Args()
	if len(args) > 0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}

	return cmd
}

func printUsage() {
	fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args[0])
}
