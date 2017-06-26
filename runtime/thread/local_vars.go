package thread

type LocalVars []Slot

func newLocalVars(maxLocals uint) LocalVars {
	return make([]Slot, maxLocals)
}

func (l LocalVars) SetInt(index int, val int32) {
	l[index].num = val
}

func (l LocalVars) GetInt(index int) int32 {
	return l[index].num
}
