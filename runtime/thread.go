package runtime

type Thread struct {
	pc    int // length of int is platform dependent
	stack *Stack
}

func NewThread() *Thread {
	return &Thread{0, newStack()}
}

func (t *Thread) Stack() *Stack {
	return t.stack
}

func (t *Thread) GetPC() int {
	return t.pc
}

func (t *Thread) SetPC(pc int) {
	t.pc = pc
}
