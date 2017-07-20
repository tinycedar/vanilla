package references

import (
	"github.com/tinycedar/vanilla/instructions"
)

func FetchInstruction(opCode uint8, reader *instructions.CodeReader) instructions.Instruction {
	switch opCode {
	case 0xb2:
		return &getstatic{opCode, reader.ReadUint16()}
	case 0xb3:
	case 0xb4:
	case 0xb5:
	case 0xb6:
		return &invokevirtual{opCode, reader.ReadUint16()}
	case 0xb7:
	case 0xb8:
		return &invokestatic{opCode, reader.ReadUint16()}
	case 0xb9:
		//TODO
	case 0xba:
		//TODO
	case 0xbb:
	case 0xbc:
		//TODO
	case 0xbd:
	case 0xbe:
		//TODO
	case 0xbf:
		//TODO
	case 0xc0:
	case 0xc1:
	case 0xc2:
		//TODO
	case 0xc3:
		//TODO
	default:
		return instructions.NewNotImplemented(opCode)
	}
	return nil
}
