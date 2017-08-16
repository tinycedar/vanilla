package conversions

import (
	"github.com/tinycedar/vanilla/instructions"
)

func FetchInstruction(opCode uint8, reader *instructions.CodeReader) instructions.Instruction {
	switch {
	case opCode == 0x85 || opCode == 0x86 || opCode == 0x87:
		return &i2x{opCode}
	case opCode == 0x91 || opCode == 0x92 || opCode == 0x93:
		return &i2x{opCode}
	default:
		return instructions.NewNotImplemented(opCode)
	}
}
