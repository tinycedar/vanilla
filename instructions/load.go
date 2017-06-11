package instructions

import (
	"fmt"
	"github.com/tinycedar/vanilla/runtime"
)

type iload_1 struct {
	opCode uint8
}

func (i *iload_1) Execute(f *runtime.Frame) {
	val := f.LocalVars().GetInt(1)
	f.OperandStack().PushInt(val)
}

func (i *iload_1) String() string {
	return fmt.Sprintf("{opcode: 0x%x, iload_1}", i.opCode)
}

type iload_2 struct {
	opCode uint8
}

func (i *iload_2) Execute(f *runtime.Frame) {
	val := f.LocalVars().GetInt(2)
	f.OperandStack().PushInt(val)
}

func (i *iload_2) String() string {
	return fmt.Sprintf("{opcode: 0x%x, iload_2}", i.opCode)
}
