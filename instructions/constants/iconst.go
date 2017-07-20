package constants

import (
	"fmt"
	"github.com/tinycedar/vanilla/runtime/thread"
)

type iconst_0 struct {
	opCode uint8
}

type iconst_1 struct {
	opCode uint8
}

type iconst_2 struct {
	opCode uint8
}

func (i *iconst_0) Execute(f *thread.Frame) {
	f.OperandStack().PushInt(0)
}

func (i *iconst_0) String() string {
	return fmt.Sprintf("{opcode: 0x%x, iconst_0}", i.opCode)
}

func (i *iconst_1) Execute(f *thread.Frame) {
	f.OperandStack().PushInt(1)
}

func (i *iconst_1) String() string {
	return fmt.Sprintf("{opcode: 0x%x, iconst_1}", i.opCode)
}

func (i *iconst_2) Execute(f *thread.Frame) {
	f.OperandStack().PushInt(2)
}

func (i *iconst_2) String() string {
	return fmt.Sprintf("{opcode: 0x%x, iconst_2}", i.opCode)
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
