package classfile

import "fmt"
import "unicode/utf16"

type ConstantUtf8Info struct {
	str string
}

func (self *ConstantUtf8Info) readInfo(reader *ClassReader) {
	length := uint32(reader.readUint16())
	bytes := reader.readBytes(length)
	self.str = decodeMUTF8(bytes)
}

//完整版本的
func decodeMUTF8(bytes []byte) string {
	utflen := len(bytes)
	chararr := make([]uint16, utflen)

	var c1, c2, c3 uint16
	count := 0
	chararr_count := 0

	for count < utflen {
		c1 = uint16(bytes[count])
		if c1 > 127 {
			break
		}
		count++
		chararr[chararr_count] = c1
		chararr_count++
	}

	for count < utflen {
		c1 = uint16(bytes[count])
		switch c1 >> 4 {
		case 0, 1, 2, 3, 4, 5, 6, 7:
			count++
			chararr[chararr_count] = c1
			chararr_count++
		case 12, 13:
			count += 2
			if count > utflen {
				panic("malformed input: partial character at end!")
			}
			c2 = uint16(bytes[count-1])
			if c2&0xC0 != 0x80 {
				panic(fmt.Errorf("malformed input around byte %v!", count))
			}
			chararr[chararr_count] = c1&0x1F<<6 | c2&0x3F
			chararr_count++
		case 14:
			count += 3
			if count > utflen {
				panic("malformed input: partial character at end!")
			}
			c2 = uint16(bytes[count-2])
			c3 = uint16(bytes[count-1])
			if c2&0xC0 != 0x80 || c3&0xC0 != 0x80 {
				panic(fmt.Errorf("malformed input around byte %v", (count - 1)))
			}
			chararr[chararr_count] = c1&0x0F<<12 | c2&0x3F<<6 | c3&0x3F<<0
			chararr_count++
		default:
			panic(fmt.Errorf("malformed input around byte %v", count))
		}
	}
	chararr = chararr[0:chararr_count]
	runes := utf16.Decode(chararr)
	return string(runes)
}
