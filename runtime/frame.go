package runtime

type Frame struct {
	next         *Frame
	localVars    LocalVars
	operandStack *OperandStack
}

func newFrame(maxLocals, maxStack uint) *Frame {
	return &Frame{nil, newLocalVars(maxLocals), newOperandStack(maxStack)}
}

type LocalVars []Slot

func newLocalVars(maxLocals uint) LocalVars {
	if maxLocals > 0 {
		return make([]Slot, maxLocals)
	}
	return nil
}

type OperandStack struct {
	size  uint
	slots []Slot
}

func newOperandStack(maxStack uint) *OperandStack {
	if maxStack > 0 {
		return &OperandStack{
			slots: make([]Slot, maxStack),
		}
	}
	return nil
}

type Slot struct {
	num int32
	ref *Object
}

type Object struct {
}
