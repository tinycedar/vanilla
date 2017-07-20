package stores

import (
	"fmt"
	"github.com/tinycedar/vanilla/runtime/thread"
)

type istore_x struct {
	opCode uint8
	index  uint8
}

func (i *istore_x) Execute(f *thread.Frame) {
	val := f.OperandStack().PopInt()
	f.LocalVars().SetInt(i.index, val)
}

func (i *istore_x) String() string {
	return fmt.Sprintf("{opcode: 0x%x, istore_%d}", i.opCode, i.index)
}
