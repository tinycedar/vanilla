package conversions

import (
	"fmt"
	"github.com/tinycedar/vanilla/runtime/thread"
)

type i2x struct {
	opCode uint8
}

func (i *i2x) Execute(f *thread.Frame) {
	stack := f.OperandStack()
	intVal := stack.PopInt()

	switch i.opCode {
	case 0x85:
		// i2l
		stack.PushLong(int64(intVal))
	case 0x86:
		// i2f
		stack.PushFloat(float32(intVal))
	case 0x87:
		// i2d
		stack.PushDouble(float64(intVal))
	case 0x91:
		// i2b: int -> byte
		stack.PushInt(int32(int8(intVal)))
	case 0x92:
		// i2c: int -> char
		stack.PushInt(int32(uint16(intVal)))
	case 0x93:
		// i2s: int -> short
		stack.PushInt(int32(int16(intVal)))
	}
}

func (i *i2x) String() string {
	return fmt.Sprintf("{opcode: 0x%x, i2x}", i.opCode)
}
