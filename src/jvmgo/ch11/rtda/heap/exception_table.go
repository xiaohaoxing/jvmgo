package heap

import "jvmgo/ch11/classfile"

type ExceptionTable []*ExceptionHandler

type ExceptionHandler struct {
	startPc   int
	endPc     int
	handlerPc int
	catchType *ClassRef
}

// 把 class 文件中的异常处理表转化成 ExceptionTable 类型
func newExceptionTable(entries []*classfile.ExceptionTableEntry, cp *ConstantPool) ExceptionTable {
	table := make([]*ExceptionHandler, len(entries))
	for i, entry := range entries {
		table[i] = &ExceptionHandler{
			startPc:   int(entry.StartPc()),
			endPc:     int(entry.EndPc()),
			handlerPc: int(entry.HandlerPc()),
			catchType: getCatchType(uint(entry.CatchType()), cp),
		}
	}
	return table
}

func getCatchType(index uint, cp *ConstantPool) *ClassRef {
	if index == 0 {
		return nil
	}
	return cp.GetConstant(index).(*ClassRef)
}

func (self ExceptionTable) findExceptionHandler(exClass *Class, pc int) *ExceptionHandler {
	for _, handler := range self {
		if pc >= handler.startPc && pc < handler.endPc {
			if handler.catchType == nil {
				return handler // catch all(为了实现 finally 子句)
			}
			catchClass := handler.catchType.ResolvedClass()
			if catchClass == exClass || catchClass.IsSuperClassOf(exClass) {
				return handler
			}
		}
	}
	return nil
}
