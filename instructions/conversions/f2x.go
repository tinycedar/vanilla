package conversions

import (
	"fmt"
	"github.com/tinycedar/vanilla/runtime/thread"
)

type f2x struct {
	opCode uint8
}

func (i *f2x) Execute(f *thread.Frame) {
	stack := f.OperandStack()
	floatVal := stack.PopFloat()

	switch i.opCode {
	case 0x8b:
		// f2i
		stack.PushInt(int32(floatVal))
	case 0x8c:
		// f2l
		stack.PushLong(int64(floatVal))
	case 0x8d:
		// f2d
		stack.PushDouble(float64(floatVal))
	}
}

func (i *f2x) String() string {
	return fmt.Sprintf("{opcode: 0x%x, f2x}", i.opCode)
}
