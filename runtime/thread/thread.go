package thread

type Thread struct {
	pc    int // length of int is platform dependent
	stack *Stack
}

func NewThread() *Thread {
	return &Thread{0, newStack()}
}

func (t *Thread) CurrentFrame() *Frame {
	return t.stack.top
}

func (t *Thread) Push(f *Frame) {
	t.stack.Push(f)
}

func (t *Thread) Pop() {
	t.stack.Pop()
}

func (t *Thread) StackNotEmpty() bool {
	return t.stack.size > 0
}
