package instructions

import (
	"fmt"
	"github.com/tinycedar/vanilla/runtime"
)

type iconst_1 struct {
	opCode uint8
}

type iconst_2 struct {
	opCode uint8
}

func (i *iconst_1) Execute(f *runtime.Frame) {
	f.OperandStack().PushInt(1)
}

func (i *iconst_1) String() string {
	return fmt.Sprintf("{opcode: 0x%x, iconst_1}", i.opCode)
}

func (i *iconst_2) Execute(f *runtime.Frame) {
	f.OperandStack().PushInt(2)
}

func (i *iconst_2) String() string {
	return fmt.Sprintf("{opcode: 0x%x, iconst_2}", i.opCode)
}
