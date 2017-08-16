package conversions

import (
	"github.com/tinycedar/vanilla/instructions"
)

func FetchInstruction(opCode uint8, reader *instructions.CodeReader) instructions.Instruction {
	switch {
	case opCode == 0x85 || opCode == 0x86 || opCode == 0x87:
		return &i2x{opCode}
	case opCode == 0x88 || opCode == 0x89 || opCode == 0x8a:
		return &l2x{opCode}
	case opCode == 0x8b || opCode == 0x8c || opCode == 0x8d:
		return &f2x{opCode}
	case opCode == 0x8e || opCode == 0x8f || opCode == 0x90:
		return &d2x{opCode}
	case opCode == 0x91 || opCode == 0x92 || opCode == 0x93:
		return &i2x{opCode}
	default:
		return instructions.NewNotImplemented(opCode)
	}
}
