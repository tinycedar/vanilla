package control

import (
	"github.com/tinycedar/vanilla/instructions"
)

func FetchInstruction(opCode uint8, reader *instructions.CodeReader) instructions.Instruction {
	switch opCode {
	case 0xa7:
		return &_goto{opCode, reader.ReadInt16()}
	case 0xac:
		return &ireturn{opCode}
	case 0xb1:
		return &_return{opCode}
	default:
		return instructions.NewNotImplemented(opCode)
	}
}
