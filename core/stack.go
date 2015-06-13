package core

/**
    A stack without thread-safe guarantee
**/
type Stack struct {
	size int
	top  *frame
}

type frame struct {
	value interface{}
	next  *frame
}

func (self *stack) Push(value *frame) bool {
	seft.top = &frame{value, self.top}
	s.size++
	return true
}

func (self *stack) Pop() (value interface{}) {
	frame := self.top
	if frame != nil {
		self.top = frame.next
		value = frame.value
		selft.size--
		return
	}
	return nil
}

func (seft *stack) Size() int {
	return self.size
}
