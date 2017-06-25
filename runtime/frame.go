package runtime

import (
	"bytes"
	"fmt"
	"github.com/tinycedar/vanilla/runtime/heap"
)

type Frame struct {
	next         *Frame
	localVars    LocalVars
	operandStack *OperandStack
	thread       *Thread
	method       *heap.Method
	nextPC       int
}

func NewFrame(t *Thread, m *heap.Method) *Frame {
	code := m.Code
	return &Frame{nil, newLocalVars(uint(code.MaxLocals)),
		newOperandStack(uint(code.MaxStack)), t, m, 0}
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

func (f *Frame) NextPC() int {
	return f.nextPC
}

func (f *Frame) SetNextPC(nextPC int) {
	f.nextPC = nextPC
}

func (f *Frame) String() string {
	var buf bytes.Buffer
	if len := len(f.localVars); len > 0 {
		buf.WriteString("LocalVars: [")
		for i, slot := range f.localVars {
			buf.WriteString(fmt.Sprintf("%d", slot.num))
			if i < len-1 {
				buf.WriteString(", ")
			}
		}
		buf.WriteString("]")
	}
	if len := f.operandStack.size; len > 0 {
		buf.WriteString("\tOperandStack: [")
		for i := 0; i < f.operandStack.size; i++ {
			slot := f.operandStack.slots[i]
			buf.WriteString(fmt.Sprintf("%d", slot.num))
			if i < len-1 {
				buf.WriteString(" <- ")
			}
		}
		buf.WriteString("]")
	}
	return buf.String()
}
