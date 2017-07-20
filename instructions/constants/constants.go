package constants

import (
	"github.com/tinycedar/vanilla/instructions"
)

func FetchInstruction(opCode uint8, reader *instructions.CodeReader) instructions.Instruction {
	switch opCode {
	case 0x03:
		return &iconst_0{opCode}
	case 0x04:
		return &iconst_1{opCode}
	case 0x05:
		return &iconst_2{opCode}
	case 0x10:
		return &bipush{opCode, reader.ReadInt8()}
	case 0x12:
		return &ldc{opCode, reader.ReadUint8()}
	default:
		return instructions.NewNotImplemented(opCode)
	}
}
