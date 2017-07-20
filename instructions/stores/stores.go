package stores

import (
	"github.com/tinycedar/vanilla/instructions"
)

func FetchInstruction(opCode uint8, reader *instructions.CodeReader) instructions.Instruction {
	switch opCode {
	case 0x3b:
		return &istore_0{opCode}
	case 0x3c:
		return &istore_1{opCode}
	case 0x3d:
		return &istore_2{opCode}
	default:
		return instructions.NewNotImplemented(opCode)
	}
}
