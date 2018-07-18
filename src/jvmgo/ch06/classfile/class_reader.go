package classfile

import "encoding/binary"
import "fmt"

type ClassReader struct {
	data []byte
}

func (self *ClassReader) readUint8() uint8 {
	val := self.data[0]
	self.data = self.data[1:]
	fmt.Printf("read a uint8:%v\n", val)
	return val
}

func (self *ClassReader) readUint16() uint16 {
	val := binary.BigEndian.Uint16(self.data)
	self.data = self.data[2:]
	fmt.Printf("read a uint16:%v\n", val)
	return val
}

func (self *ClassReader) readUint32() uint32 {
	val := binary.BigEndian.Uint32(self.data)
	self.data = self.data[4:]
	fmt.Printf("read a uint32:%v\n", val)
	return val
}

func (self *ClassReader) readUint64() uint64 {
	val := binary.BigEndian.Uint64(self.data)
	self.data = self.data[8:]
	fmt.Printf("read a uint64:%v\n", val)
	return val
}
//第一个字节表示需要读取的大小。
func (self *ClassReader) readUint16s() []uint16 {
	n := self.readUint16()
	s := make([]uint16, n)	//大小为n的uint数组
	for i := range s {
		s[i] = self.readUint16()
	}
	return s
}
//读取指定数量的字节
func (self *ClassReader) readBytes(length uint32) []byte {
	bytes := self.data[:length]
	self.data = self.data[length:]
	return bytes
}