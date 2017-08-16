package conversions

import (
	"fmt"
	"github.com/tinycedar/vanilla/runtime/thread"
)

type l2x struct {
	opCode uint8
}

func (i *l2x) Execute(f *thread.Frame) {
	stack := f.OperandStack()
	longVal := stack.PopLong()

	switch i.opCode {
	case 0x88:
		// l2i
		stack.PushInt(int32(longVal))
	case 0x89:
		// l2f
		stack.PushFloat(float32(longVal))
	case 0x8a:
		// l2d
		stack.PushDouble(float64(longVal))
	}
}

func (i *l2x) String() string {
	return fmt.Sprintf("{opcode: 0x%x, l2x}", i.opCode)
}
