package control

import (
	"github.com/tinycedar/vanilla/instructions"
)

func FetchInstruction(opCode uint8, reader *instructions.CodeReader) instructions.Instruction {
	switch {
	case opCode == 0xa7:
		return &_goto{opCode, reader.ReadInt16()}
	case 0xac <= opCode && opCode <= 0xaf:
		return &xreturn{opCode}
	case opCode == 0xb1:
		return &_return{opCode}
	default:
		return instructions.NewNotImplemented(opCode)
	}
}
