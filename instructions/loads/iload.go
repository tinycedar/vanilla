package loads

import (
	"fmt"
	"github.com/tinycedar/vanilla/runtime/thread"
)

type iload_0 struct {
	opCode uint8
}

type iload_1 struct {
	opCode uint8
}

func (i *iload_0) Execute(f *thread.Frame) {
	_iload(f, 0)
}

func (i *iload_0) String() string {
	return fmt.Sprintf("{opcode: 0x%x, iload_0}", i.opCode)
}

func (i *iload_1) Execute(f *thread.Frame) {
	_iload(f, 1)
}

func (i *iload_1) String() string {
	return fmt.Sprintf("{opcode: 0x%x, iload_1}", i.opCode)
}

type iload_2 struct {
	opCode uint8
}

func (i *iload_2) Execute(f *thread.Frame) {
	_iload(f, 2)
}

func (i *iload_2) String() string {
	return fmt.Sprintf("{opcode: 0x%x, iload_2}", i.opCode)
}

func _iload(f *thread.Frame, index int) {
	val := f.LocalVars().GetInt(index)
	f.OperandStack().PushInt(val)
}
