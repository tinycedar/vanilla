package comparisons

import (
	"github.com/tinycedar/vanilla/instructions"
)

func FetchInstruction(opCode uint8, reader *instructions.CodeReader) instructions.Instruction {
	switch opCode {
	case 0xa2:
		return &if_icmpge{opCode, reader.ReadUint16()}
	default:
		return instructions.NewNotImplemented(opCode)
	}
}
