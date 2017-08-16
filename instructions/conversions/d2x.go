package conversions

import (
	"fmt"
	"github.com/tinycedar/vanilla/runtime/thread"
)

type d2x struct {
	opCode uint8
}

func (i *d2x) Execute(f *thread.Frame) {
	stack := f.OperandStack()
	doubleValue := stack.PopLong()

	switch i.opCode {
	case 0x8e:
		// d2i
		stack.PushInt(int32(doubleValue))
	case 0x8f:
		// d2l
		stack.PushLong(int64(doubleValue))
	case 0x90:
		// d2f
		stack.PushFloat(float32(doubleValue))
	}
}

func (i *d2x) String() string {
	return fmt.Sprintf("{opcode: 0x%x, d2x}", i.opCode)
}
