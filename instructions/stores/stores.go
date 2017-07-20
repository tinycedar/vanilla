package stores

import (
	"github.com/tinycedar/vanilla/instructions"
)

func FetchInstruction(opCode uint8, reader *instructions.CodeReader) instructions.Instruction {
	switch opCode {
	case 0x36:
		return &istore_x{opCode, reader.ReadUint8()}
	case 0x3b:
		return &istore_x{opCode, 0}
	case 0x3c:
		return &istore_x{opCode, 1}
	case 0x3d:
		return &istore_x{opCode, 2}
	case 0x3e:
		return &istore_x{opCode, 3}
	default:
		return instructions.NewNotImplemented(opCode)
	}
}
