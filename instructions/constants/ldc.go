package constants

import (
	"fmt"
	"github.com/tinycedar/classp/classfile"
	"github.com/tinycedar/vanilla/runtime/thread"
)

type ldc struct {
	opCode uint8
	index  uint8
}

func (i *ldc) Execute(f *thread.Frame) {
	_ldc(f, i.index)
}

func _ldc(frame *thread.Frame, index uint8) {
	stack := frame.OperandStack()
	if cp, ok := frame.Method().Cp[index].(*classfile.ConstantStringInfo); ok {
		stack.PushString(cp.String(frame.Method().Cp))
	}
	//TODO
}

func (i *ldc) String() string {
	return fmt.Sprintf("{opcode: 0x%x, ldc}", i.opCode)
}
