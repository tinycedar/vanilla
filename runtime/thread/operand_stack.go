package thread

import (
	"fmt"
	"math"
)

type OperandStack struct {
	size  int
	slots []Slot
}

func newOperandStack(maxStack uint) *OperandStack {
	return &OperandStack{
		slots: make([]Slot, maxStack),
	}
}

func (s *OperandStack) PushInt(val int32) {
	s.slots[s.size].num = val
	s.size++
}

func (s *OperandStack) PopInt() int32 {
	s.size--
	return s.slots[s.size].num
}

func (s *OperandStack) PushLong(val int64) {
	s.slots[s.size].num = int32(val >> 32) // high
	s.slots[s.size+1].num = int32(val)     // low
	s.size += 2
}

func (s *OperandStack) PopLong() int64 {
	s.size -= 2
	high := s.slots[s.size].num
	low := s.slots[s.size+1].num
	return int64(high)<<32 | int64(low)
}

func (s *OperandStack) PushFloat(val float32) {
	s.slots[s.size].num = int32(math.Float32bits(val))
	s.size++
}

func (s *OperandStack) PopFloat() float32 {
	s.size--
	return math.Float32frombits(uint32(s.slots[s.size].num))
}

func (s *OperandStack) PushDouble(val float64) {
	s.slots[s.size].num = int32(math.Float64bits(val))
	s.size++
}

func (s *OperandStack) PopDouble() float64 {
	s.size--
	return math.Float64frombits(uint64(s.slots[s.size].num))
}

func (s *OperandStack) PushString(str string) {
	s.slots[s.size] = Slot{ref: &Object{str}}
	s.size++
}

func (s *OperandStack) PopString() string {
	s.size--
	return fmt.Sprintf("%s", s.slots[s.size].ref.data)
}

func (s *OperandStack) PushRef(ref *Object) {
	s.slots[s.size].ref = ref
	s.size++
}

func (s *OperandStack) PopRef() *Object {
	s.size--
	return s.slots[s.size].ref
}

func (s *OperandStack) PushSlot(slot Slot) {
	s.slots[s.size] = slot
	s.size++
}
func (s *OperandStack) PopSlot() Slot {
	s.size--
	return s.slots[s.size]
}
