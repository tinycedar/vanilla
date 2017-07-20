package comparisons

import (
	"github.com/tinycedar/vanilla/instructions"
)

func FetchInstruction(opCode uint8, reader *instructions.CodeReader) instructions.Instruction {
	switch {
	case 0x99 <= opCode && opCode <= 0x9e:
		return &if_cond{opCode, reader.ReadInt16()}
	case 0x9f <= opCode && opCode <= 0xa4:
		return &if_icmp_cond{opCode, reader.ReadInt16()}
	case 0xa5 <= opCode && opCode <= 0xa6:
		return &if_acmp_cond{opCode, reader.ReadInt16()}
	default:
		return instructions.NewNotImplemented(opCode)
	}
}
