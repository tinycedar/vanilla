package runtime

type Frame struct {
	next         *Frame
	localVars    []Slot
	operandStack *OperandStack
}

type OperandStack struct {
	size  uint
	slots []Slot
}

type Slot struct {
	num int32
	ref *Object
}

type Object struct {
}

func newFrame(maxLocals, maxStack uint16) *Frame {
	return &Frame{nil, make([]Slot, maxLocals), &OperandStack{
		slots: make([]Slot, maxStack),
	}}
}
