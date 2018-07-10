package classpath

import "os"
import "path/filepath"

type Classpath struct {
	 bootClasspath Entry
	 extClasspath Entry
	 userClasspath Entry
}
//使用-Xjre选项解析“启动类”和“拓展类”路径。
//使用-cp/-classpath选项解析“用户类”路径
func Parse(jreOption, cpOption string) *Classpath{
	cp := &Classpath{}
	cp.parseBootAndExtClasspath(jreOption)
	cp.parseUserClasspath(cpOption)
	return cp
}

//依次在“启动类”，“拓展类”，“用户类”路径的顺序搜索class文件并读取
func (self *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	className  = className + ".class"

	if data, entry, err := self.bootClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	
	if data, entry, err := self.extClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	return self.userClasspath.readClass(className)
}

func (self *Classpath) String() string{
	return self.userClasspath.String()
}

/*
 private methods
*/
//读取“启动类”和“拓展类”路径
func (self *Classpath) parseBootAndExtClasspath(jreOption string){
	//先确认真实的路径，按照-Xjre选项、当前目录、JAVA_HOME环境变量的顺序试
	jreDir := getJreDir(jreOption)

	//	jre/lib/*
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	self.bootClasspath = newWildcardEntry(jreLibPath)

	//  jre/lib/ext/*
	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	self.extClasspath = newWildcardEntry(jreExtPath)
}

//优先使用-Xjre选项作为jre目录
//然后当前目录
//然后JAVA_HOME环境变量
func getJreDir(jreOption string) string {
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}
	if exists("./jre") {
		return "./jre"
	}
	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre")
	}
	panic("Can not find jre folder!")
}

//判断路径是否存在
func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

//处理“用户类”路径，如果没有设置-cp选项则使用当前目录
func (self *Classpath) parseUserClasspath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}
	self.userClasspath = newEntry(cpOption)
}

