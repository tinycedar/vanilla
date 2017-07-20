package constants

import (
	"github.com/tinycedar/vanilla/instructions"
	"github.com/tinycedar/vanilla/runtime/thread"
)

type nop struct {
	opCode uint8
}

func (n *nop) Execute(f *thread.Frame) {}

func FetchInstruction(opCode uint8, reader *instructions.CodeReader) instructions.Instruction {
	switch opCode {
	case 0x00:
		return &nop{opCode}
	case 0x01:
		return &aconst_null{opCode}
	case 0x02:
		return &iconst_x{opCode, -1}
	case 0x03:
		return &iconst_x{opCode, 0}
	case 0x04:
		return &iconst_x{opCode, 1}
	case 0x05:
		return &iconst_x{opCode, 2}
	case 0x06:
		return &iconst_x{opCode, 3}
	case 0x07:
		return &iconst_x{opCode, 4}
	case 0x08:
		return &iconst_x{opCode, 5}
		//TODO
	case 0x10:
		return &bipush{opCode, reader.ReadInt8()}
	case 0x12:
		return &ldc{opCode, reader.ReadUint8()}
	default:
		return instructions.NewNotImplemented(opCode)
	}
}
