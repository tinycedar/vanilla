package constants

import (
	"fmt"
	"github.com/tinycedar/vanilla/runtime/thread"
)

type bipush struct {
	opCode uint8
	value  int8
}

type sipush struct {
	opCode uint8
	value  int16
}

func (i *bipush) Execute(f *thread.Frame) {
	f.OperandStack().PushInt(int32(i.value))
}

func (i *sipush) Execute(f *thread.Frame) {
	f.OperandStack().PushInt(int32(i.value))
}

func (i *bipush) String() string {
	return fmt.Sprintf("{opcode: 0x%x, bipush}", i.opCode)
}

func (i *sipush) String() string {
	return fmt.Sprintf("{opcode: 0x%x, sipush}", i.opCode)
}
