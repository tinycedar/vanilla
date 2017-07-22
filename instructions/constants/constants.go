package constants

import (
	"fmt"
	"github.com/tinycedar/vanilla/instructions"
	"github.com/tinycedar/vanilla/runtime/thread"
)

type nop struct {
	opCode uint8
}

func (n *nop) Execute(f *thread.Frame) {}

func FetchInstruction(opCode uint8, reader *instructions.CodeReader) instructions.Instruction {
	switch opCode {
	case 0x00:
		return &nop{opCode}
	case 0x01:
		return &aconst_null{opCode}
	case 0x02:
		return &iconst_x{opCode, -1}
	case 0x03:
		return &iconst_x{opCode, 0}
	case 0x04:
		return &iconst_x{opCode, 1}
	case 0x05:
		return &iconst_x{opCode, 2}
	case 0x06:
		return &iconst_x{opCode, 3}
	case 0x07:
		return &iconst_x{opCode, 4}
	case 0x08:
		return &iconst_x{opCode, 5}
	case 0x09:
		return &lconst_x{opCode, 0}
	case 0x0a:
		return &lconst_x{opCode, 1}
	case 0x0b:
		return &fconst_x{opCode, 0.0}
	case 0x0c:
		return &fconst_x{opCode, 1.0}
	case 0x0d:
		return &fconst_x{opCode, 2.0}
	case 0x0e:
		return &dconst_x{opCode, 0.0}
	case 0x0f:
		return &dconst_x{opCode, 1.0}
	case 0x10:
		return &bipush{opCode, reader.ReadInt8()}
	case 0x11:
		return &sipush{opCode, reader.ReadInt16()}
	case 0x12:
		return &ldc_x{opCode, uint16(reader.ReadUint8())}
	case 0x13:
		fallthrough
	case 0x14:
		return &ldc_x{opCode, reader.ReadUint16()}
	default:
		panic(fmt.Sprintf("invalid opCode: %d", opCode))
	}
}
