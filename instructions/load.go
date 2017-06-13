package instructions

import (
	"fmt"
	"github.com/tinycedar/vanilla/runtime"
)

type iload_1 struct {
	opCode uint8
}

func (i *iload_1) Execute(f *runtime.Frame) {
	_iload(f, 1)
}

func (i *iload_1) String() string {
	return fmt.Sprintf("{opcode: 0x%x, iload_1}", i.opCode)
}

type iload_2 struct {
	opCode uint8
}

func (i *iload_2) Execute(f *runtime.Frame) {
	_iload(f, 2)
}

func (i *iload_2) String() string {
	return fmt.Sprintf("{opcode: 0x%x, iload_2}", i.opCode)
}

func _iload(f *runtime.Frame, index int) {
	val := f.LocalVars().GetInt(index)
	f.OperandStack().PushInt(val)
}
