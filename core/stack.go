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

func (self *Stack) Push(value *frame) bool {
	self.top = &frame{value, self.top}
	self.size++
	return true
}

func (self *Stack) Pop() (value interface{}) {
	frame := self.top
	if frame != nil {
		self.top = frame.next
		value = frame.value
		self.size--
		return
	}
	return nil
}

func (self *Stack) Size() int {
	return self.size
}
