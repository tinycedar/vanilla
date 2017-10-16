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
	cpInfo := f.Method().Cp[i.index]
	if cp, ok := cpInfo.(*classfile.ConstantStringInfo); ok {
		f.OperandStack().PushString(cp.String(f.Method().Cp))
	}
	if cp, ok := cpInfo.(*classfile.ConstantFloatInfo); ok {
		f.OperandStack().PushFloat(cp.Value())
	}
	//TODO
}

func (i *ldc_x) String() string {
	return fmt.Sprintf("{opcode: 0x%x, ldc_x}", i.opCode)
}
