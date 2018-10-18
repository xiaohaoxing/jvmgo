/*
	定义了DeprecatedAttribute和SyntheticAttribute2种属性结构体。
	他们不包含任何信息，因此readInfo函数是空的。
*/
package classfile

type DeprecatedAttribute struct{ MarkerAttribute }
type SyntheticAttribute struct{ MarkerAttribute }

type MarkerAttribute struct{}

func (self *MarkerAttribute) readInfo(reader *ClassReader) {
	//read nothing
}
