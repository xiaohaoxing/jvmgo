package classfile

type ConstantStringInfo struct {
	cp ConstantPool
	stringIndex uint16
}

func (self *ConstantStringInfo) readInfo(reader *ClassReader) {

}

func (self *ConstantStringInfo) String() string {
	return self.cp.getUtf8(self.stringIndex)
}