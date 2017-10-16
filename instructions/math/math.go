package math

import (
	"github.com/tinycedar/vanilla/instructions"
)

func FetchInstruction(opCode uint8, reader *instructions.CodeReader) instructions.Instruction {
	switch {
	case 0x60 <= opCode && opCode <= 0x63:
		return &xadd{opCode}
	case opCode == 0x84:
		return &iinc{opCode, reader.ReadUint8(), reader.ReadInt8()}
	default:
		return instructions.NewNotImplemented(opCode)
	}
}
