package classfile

type AttributeInfo interface {
	readInfo(reader *ClassReader)
}

func readAttributes(reader *ClassReader, cp ConstantPool) []AttributeInfo {
	attributesCount := reader.readUint16()
	attributes := make([]AttributeInfo, attributesCount)
	for i := range attributes {
		attributes[i] = readAttribute(reader, cp)
	}
	return attributes
}

/*
	根据属性名索引在常量池中找到属性名，然后读取属性长度。创建属性对象。
*/
func readAttribute(reader *ClassReader, cp ConstantPool) AttributeInfo {
	attrNameIndex := reader.readUint16()
	attrName := cp.getUtf8(attrNameIndex)
	attrLen := reader.readUint32()
	attrInfo := newAttributeInfo(attrName, attrLen, cp)
	attrInfo.readInfo(reader)
	return attrInfo
}

/*
	解析23种JVM规范里面的属性。先做其中8种。
*/
func newAttributeInfo(attrName string, attrLen uint32, cp ConstantPool) AttributeInfo {
	switch attrName {
	//case "AnnotationDefault":
	case "BootstrapMethods":
		return &BootstrapMethodsAttribute{}
	case "Code":
		return &CodeAttribute{cp: cp}
	case "ConstantValue":
		return &ConstantValueAttribute{}
	case "Deprecated":
		return &DeprecatedAttribute{}
	case "EnclosingMethod":
		return &EnclosingMethodAttribute{cp: cp}
	case "Exceptions":
		return &ExceptionsAttribute{}
	case "InnerClasses":
		return &InnerClassesAttribute{}
	case "LineNumberTable":
		return &LineNumberTableAttribute{}
	case "LocalVariableTable":
		return &LocalVariableTableAttribute{}
	case "LocalVariableTypeTable":
		return &LocalVariableTypeTableAttribute{}
	//case "MethodParameters":
	//case "RuntimeInvisibleAnnotations":
	//case "RuntimeInvisibleParameterAnnotations":
	//case "RuntimeInvisibleTypeAnnotations":
	//case "RuntimeVisibleAnnotations":
	//case "RuntimeVisibleParameterAnnotations":
	//case "RuntimeVisibleTypeAnnotations":
	case "Signature":
		return &SignatureAttribute{cp: cp}
	case "SourceFile":
		return &SourceFileAttribute{}
	//case "SourceDebugExtension":
	//case "StackMapTable":
	case "Synthetic":
		return &SyntheticAttribute{}
	default:
		return &UnparsedAttribute{attrName, attrLen, nil}
	}
}
