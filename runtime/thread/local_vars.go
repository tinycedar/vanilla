package thread

type LocalVars []Slot

func newLocalVars(maxLocals uint) LocalVars {
	return make([]Slot, maxLocals)
}

func (l LocalVars) SetInt(index uint8, val int32) {
	l[index].num = val
}

func (l LocalVars) GetInt(index uint8) int32 {
	return l[index].num
}

func (l LocalVars) SetLong(index uint8, val int64) {
	// big endian
	high := int32(val >> 32)
	low := int32(val)
	l[index].num = high
	l[index+1].num = low
}

func (l LocalVars) GetLong(index uint8) int64 {
	high := l[index].num
	low := l[index+1].num
	return int64(high<<32) | int64(low)
}

func (l LocalVars) SetSlot(index uint, slot Slot) {
	l[index] = slot
}
