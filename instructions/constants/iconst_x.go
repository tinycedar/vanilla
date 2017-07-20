package constants

import (
	"fmt"
	"github.com/tinycedar/vanilla/runtime/thread"
)

type iconst_x struct {
	opCode   uint8
	constVal int32
}

func (i *iconst_x) Execute(f *thread.Frame) {
	f.OperandStack().PushInt(i.constVal)
}

func (i *iconst_x) String() string {
	return fmt.Sprintf("{opcode: 0x%x, iconst_%d}", i.opCode, i.constVal)
}

type bipush struct {
	opCode uint8
	value  int8
}

func (i *bipush) Execute(f *thread.Frame) {
	f.OperandStack().PushInt(int32(i.value))
}

func (i *bipush) String() string {
	return fmt.Sprintf("{opcode: 0x%x, bipush}", i.opCode)
}
