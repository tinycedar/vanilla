package loads

import (
	"github.com/tinycedar/vanilla/instructions"
)

func FetchInstruction(opCode uint8, reader *instructions.CodeReader) instructions.Instruction {
	switch opCode {
	case 0x15:
		return &iload_x{opCode, reader.ReadUint8()}
	case 0x1a:
		return &iload_x{opCode, 0}
	case 0x1b:
		return &iload_x{opCode, 1}
	case 0x1c:
		return &iload_x{opCode, 2}
	case 0x1d:
		return &iload_x{opCode, 3}
	default:
		return instructions.NewNotImplemented(opCode)
	}
}
