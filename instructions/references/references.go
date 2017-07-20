package references

import (
	"github.com/tinycedar/vanilla/instructions"
)

func FetchInstruction(opCode uint8, reader *instructions.CodeReader) instructions.Instruction {
	switch opCode {
	case 0xb6:
		return &invokevirtual{reader.ReadUint16(), opCode}
	case 0xb8:
		return &invokestatic{reader.ReadUint16(), opCode}
	default:
		return instructions.NewNotImplemented(opCode)
	}
}
