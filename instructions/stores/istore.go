package stores

import (
	"fmt"
	"github.com/tinycedar/vanilla/runtime/thread"
)

type istore_0 struct {
	opCode uint8
}

type istore_1 struct {
	opCode uint8
}

func (i *istore_0) Execute(f *thread.Frame) {
	val := f.OperandStack().PopInt()
	f.LocalVars().SetInt(0, val)
}

func (i *istore_0) String() string {
	return fmt.Sprintf("{opcode: 0x%x, istore_0}", i.opCode)
}

func (i *istore_1) Execute(f *thread.Frame) {
	val := f.OperandStack().PopInt()
	f.LocalVars().SetInt(1, val)
}

func (i *istore_1) String() string {
	return fmt.Sprintf("{opcode: 0x%x, istore_1}", i.opCode)
}

type istore_2 struct {
	opCode uint8
}

func (i *istore_2) Execute(f *thread.Frame) {
	val := f.OperandStack().PopInt()
	f.LocalVars().SetInt(2, val)
}

func (i *istore_2) String() string {
	return fmt.Sprintf("{opcode: 0x%x, istore_2}", i.opCode)
}
