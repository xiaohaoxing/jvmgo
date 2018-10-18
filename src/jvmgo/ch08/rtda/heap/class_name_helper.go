package heap

func getArrayClassName(className string) string {
	return "[" + toDescriptor(className)
}

func toDescriptor(className string) string {
	if className[0] == '[' {
		return className
	}
	if d, ok := primitiveTypes[className]; ok {
		return d
	}
	return "L" + className + ";"
}

var primitiveTypes = map[string]string{
	"void":    "V",
	"boolean": "Z",
	"byte":    "B",
	"short":   "S",
	"int":     "I",
	"long":    "L",
	"char":    "C",
	"float":   "F",
	"double":  "D",
}

func getComponentClassName(className string) string {
	if className[0] == '[' {
		componentTypeDescriptor := className[1:]
		return toClassName(componentTypeDescriptor)
	}
	panic("Not array: " + className)
}

//从类描述中得到类名
func toClassName(descriptor string) string {
	// a array：去掉前面的括号就是类名
	if descriptor[0] == '[' {
		return descriptor
	}
	// object：去掉前面的 L 和最后的分号就是类名
	if descriptor[0] == 'L' {
		return descriptor[1 : len(descriptor)-1]
	}
	// 从基本类型的关系中找到类名
	for className, d := range primitiveTypes {
		if d == descriptor {
			return className
		}
	}
	panic("Invalid descriptor: " + descriptor)
}
