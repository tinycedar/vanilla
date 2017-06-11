package instructions

import (
	"fmt"
	"github.com/tinycedar/class-parser/classfile"
	"github.com/tinycedar/vanilla/runtime"
)

type ldc struct {
	opCode uint8
	index  uint8
}

func (i *ldc) Execute(f *runtime.Frame) {
	_ldc(f, i.index)
}

func _ldc(frame *runtime.Frame, index uint8) {
	stack := frame.OperandStack()
	if cp, ok := frame.Method().ConstantPool()[index].(*classfile.ConstantStringInfo); ok {
		stack.PushString(cp.String(frame.Method().ConstantPool()))
	}
	//TODO
}

func (i *ldc) String() string {
	return fmt.Sprintf("{opcode: 0x%x, ldc}", i.opCode)
}