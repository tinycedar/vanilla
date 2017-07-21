package thread

import "github.com/tinycedar/vanilla/runtime/heap"

type Thread struct {
	stack *Stack
}

func NewThread() *Thread {
	return &Thread{newStack()}
}

func (t *Thread) NewFrame(m *heap.Method) *Frame {
	frame := newFrame(t, m)
	t.push(frame)
	return frame
}

func (t *Thread) CurrentFrame() *Frame {
	return t.stack.top
}

func (t *Thread) push(f *Frame) {
	t.stack.Push(f)
}

func (t *Thread) Pop() {
	t.stack.Pop()
}

func (t *Thread) StackNotEmpty() bool {
	return t.stack.size > 0
}
