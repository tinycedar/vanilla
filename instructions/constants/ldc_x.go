package constants

import (
	"fmt"
	"github.com/tinycedar/classp/classfile"
	"github.com/tinycedar/vanilla/runtime/thread"
)

type ldc_x struct {
	opCode uint8
	index  uint16
}

func (i *ldc_x) Execute(f *thread.Frame) {
	if cp, ok := f.Method().Cp[i.index].(*classfile.ConstantStringInfo); ok {
		f.OperandStack().PushString(cp.String(f.Method().Cp))
	}
	//TODO
}

func (i *ldc_x) String() string {
	return fmt.Sprintf("{opcode: 0x%x, ldc_x}", i.opCode)
}
