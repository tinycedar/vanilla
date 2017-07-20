package instructions

type CodeReader struct {
	code   []byte
	offset int
}

func NewByteCodeReader(code []byte, offset int) *CodeReader {
	return &CodeReader{code, offset}
}

func (r *CodeReader) Offset() int {
	return r.offset
}

func (r *CodeReader) ReadUint8() uint8 {
	return uint8(r.ReadBytes(1)[0])
}

func (r *CodeReader) ReadUint16() uint16 {
	byte1 := uint16(r.ReadUint8())
	byte2 := uint16(r.ReadUint8())
	return (byte1 << 8) | byte2
}

func (r *CodeReader) ReadInt8() int8 {
	return int8(r.ReadBytes(1)[0])
}

func (r *CodeReader) ReadInt16() int16 {
	byte1 := int16(r.ReadUint8())
	byte2 := int16(r.ReadUint8())
	return (byte1 << 8) | byte2
}

func (r *CodeReader) ReadBytes(len int) []byte {
	result := r.code[r.offset : r.offset+len]
	r.offset += len
	return result
}
