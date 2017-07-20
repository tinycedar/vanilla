package loads

import (
	"github.com/tinycedar/vanilla/instructions"
)

func FetchInstruction(opCode uint8, reader *instructions.CodeReader) instructions.Instruction {
	switch opCode {
	case 0x1a:
		return &iload_0{opCode}
	case 0x1b:
		return &iload_1{opCode}
	case 0x1c:
		return &iload_2{opCode}
	default:
		return instructions.NewNotImplemented(opCode)
	}
}
