package classpath

import "errors"
import "strings"

type CompositeEntry []Entry

func newCompositeEntry(pathList string) CompositeEntry{
	//创建entry
	compositeEntry := []Entry{}
	//slice类型
	for _, path := range strings.Split(pathList, pathListSeparator) {
		entry := newEntry(path)
		compositeEntry = append(compositeEntry, entry)
	}
	return compositeEntry
}
//读取class文件内容
func (self CompositeEntry) readClass(className string) ([]byte, Entry, error){
	for _, entry := range self {
		data, from, err := entry.readClass(className)
		if err == nil {
			return data, from, nil
		}
	}
	return nil, nil, errors.New("Class not found: " + className)
}
//拼接每一个子路径的String方法返回的字符串，用pathListSeparator分割
func (self CompositeEntry) String() string{
	strs := make([]string, len(self))

	for i, entry := range self {
		strs[i] = entry.String()
	}

	return strings.Join(strs, pathListSeparator)
}