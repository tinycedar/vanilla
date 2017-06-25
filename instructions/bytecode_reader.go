package instructions

import (
	"encoding/binary"
)

var bigEndian = binary.BigEndian

type bytecodeReader struct {
	bytecode []byte
	index    int
}

func NewByteCodeReader(bytes []byte, index int) *bytecodeReader {
	return &bytecodeReader{bytes, index}
}

func (r *bytecodeReader) FetchInstruction() Instruction {
	opCode := r.fetchOpCode()
	//fmt.Printf("opCode: 0x%x\n", opCode)
	switch opCode {
	case 0x04:
		return &iconst_1{opCode}
	case 0x05:
		return &iconst_2{opCode}
	case 0x12:
		return &ldc{opCode, r.ReadUint8()}
	case 0x1b:
		return &iload_1{opCode}
	case 0x1c:
		return &iload_2{opCode}
	case 0x3c:
		return &istore_1{opCode}
	case 0x3d:
		return &istore_2{opCode}
	case 0x60:
		return &iadd{opCode}
	case 0xb1:
		return &_return{opCode}
	case 0xb6:
		return &invokevirtual{r.ReadUint16(), opCode}
	case 0xb8:
		return &invokestatic{r.ReadUint16(), opCode}
	default:
		//fmt.Printf("invalid opcode: %d, %x\n", opCode, opCode)
	}
	return nil
}

func (r *bytecodeReader) NextPC() int {
	return r.index
}

func (r *bytecodeReader) fetchOpCode() uint8 {
	return r.ReadUint8()
}

func (r *bytecodeReader) ReadUint8() uint8 {
	return uint8(r.ReadBytes(1)[0])
}

func (r *bytecodeReader) ReadUint16() uint16 {
	byte1 := uint16(r.ReadUint8())
	byte2 := uint16(r.ReadUint8())
	return (byte1 << 8) | byte2
}

func (r *bytecodeReader) ReadBytes(len int) []byte {
	result := r.bytecode[r.index : r.index+len]
	r.index += len
	return result
}
