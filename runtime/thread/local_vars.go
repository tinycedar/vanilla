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
