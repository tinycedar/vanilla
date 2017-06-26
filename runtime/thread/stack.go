package thread

import (
	"bytes"
	"fmt"
)

type Stack struct {
	size uint
	top  *Frame
}

func newStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(f *Frame) {
	f.next = s.top
	s.top = f
	s.size++
}

func (s *Stack) Pop() *Frame {
	f := s.top
	if f != nil {
		s.top = f.next
		s.size--
		return f
	}
	return nil
}

func (s *Stack) Size() uint {
	return s.size
}

func (s *Stack) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("Stack size: %d\n", s.size))
	for f := s.top; f != nil; f = f.next {
		buffer.WriteString(fmt.Sprintf("%v", f))
		buffer.WriteString("\n")
	}
	return buffer.String()
}
