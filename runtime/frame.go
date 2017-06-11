package runtime

import (
	"bytes"
	"fmt"
	"github.com/tinycedar/class-parser/classfile"
)

type Frame struct {
	next         *Frame
	localVars    LocalVars
	operandStack *OperandStack
	thread       *Thread
	method       *classfile.MethodInfo
}

func NewFrame(t *Thread, maxLocals, maxStack uint, m *classfile.MethodInfo) *Frame {
	return &Frame{nil, newLocalVars(maxLocals),
		newOperandStack(maxStack), t, m}
}

func (f *Frame) LocalVars() LocalVars {
	return f.localVars
}

func (f *Frame) OperandStack() *OperandStack {
	return f.operandStack
}

func (f *Frame) Method() *classfile.MethodInfo {
	return f.method
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
