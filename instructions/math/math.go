package math

import (
	"github.com/tinycedar/vanilla/instructions"
)

func FetchInstruction(opCode uint8, reader *instructions.CodeReader) instructions.Instruction {
	switch opCode {
	case 0x60:
		return &iadd{opCode}
	case 0x84:
		return &iinc{opCode, reader.ReadUint8(), reader.ReadInt8()}
	default:
		return instructions.NewNotImplemented(opCode)
	}
}
