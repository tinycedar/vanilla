package thread

import (
	"github.com/tinycedar/vanilla/runtime/heap"
)

type Frame struct {
	next         *Frame
	localVars    LocalVars
	operandStack *OperandStack
	thread       *Thread
	method       *heap.Method
	pc           int
}

func NewFrame(t *Thread, m *heap.Method) *Frame {
	code := m.CodeAttr
	return &Frame{nil, newLocalVars(uint(code.MaxLocals)),
		newOperandStack(uint(code.MaxStack)), t, m, 0}
}

func (f *Frame) NextFrame() *Frame {
	return f.next
}

func (f *Frame) LocalVars() LocalVars {
	return f.localVars
}

func (f *Frame) OperandStack() *OperandStack {
	return f.operandStack
}

func (f *Frame) Thread() *Thread {
	return f.thread
}

func (f *Frame) Method() *heap.Method {
	return f.method
}

func (f *Frame) GetPC() int {
	return f.pc
}

func (f *Frame) SetPC(pc int) {
	f.pc = pc
}
