package classfile

type ConstantStringInfo struct {
	cp ConstantPool
	nameIndex uint16
}

func (self *ConstantStringInfo) readInfo(reader *ClassReader) {
	self.nameIndex = reader.readUint16()
}

func (self *ConstantStringInfo) String() string {
	return self.cp.getUtf8(self.nameIndex)
}