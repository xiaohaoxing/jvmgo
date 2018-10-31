package heap

func (self *Object) Clone() *Object {
	return &Object{
		class: self.class,
		data:  self.cloneData(),
	}
}

func (self *Object) cloneData() interface{} {
	switch self.data.(type) {
	case []int8:
		ints := self.data.([]int8)
		ints2 := make([]int8, len(ints))
		copy(ints2, ints)
		return ints2
	case []int16:
		ints := self.data.([]int16)
		ints2 := make([]int16, len(ints))
		copy(ints2, ints)
		return ints2
	case []uint16:
		ints := self.data.([]uint16)
		ints2 := make([]uint16, len(ints))
		copy(ints2, ints)
		return ints2
	case []int32:
		ints := self.data.([]int32)
		ints2 := make([]int32, len(ints))
		copy(ints2, ints)
		return ints2
	case []int64:
		ints := self.data.([]int64)
		ints2 := make([]int64, len(ints))
		copy(ints2, ints)
		return ints2
	case []float32:
		floats := self.data.([]float32)
		floats2 := make([]float32, len(floats))
		copy(floats2, floats)
		return floats2
	case []float64:
		floats := self.data.([]float64)
		floats2 := make([]float64, len(floats))
		copy(floats2, floats)
		return floats2
	case []*Object:
		elements := self.data.([]*Object)
		elements2 := make([]*Object, len(elements))
		copy(elements2, elements)
		return elements2
	default:
		slots := self.data.(Slots)
		slots2 := newSlots(uint(len(slots)))
		copy(slots2, slots)
		return slots2
	}
}
