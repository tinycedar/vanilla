package runtime

import "fmt"

type OperandStack struct {
	size  int
	slots []Slot
}

func newOperandStack(maxStack uint) *OperandStack {
	return &OperandStack{
		slots: make([]Slot, maxStack),
	}
}

func (s *OperandStack) PushInt(num int32) {
	s.slots[s.size].num = num
	s.size++
}

func (s *OperandStack) PopInt() int32 {
	s.size--
	return s.slots[s.size].num
}

func (s *OperandStack) PushString(str string) {
	s.slots[s.size] = Slot{ref: &Object{str}}
	s.size++
}

func (s *OperandStack) PopString() string {
	s.size--
	return fmt.Sprintf("%s", s.slots[s.size].ref.data)
}
