package binary

import (
	"encoding/binary"
	"io"
)

//go:nobounds
func PutUint64(buf []byte, u uint64) error {
	if len(buf)+Uint64Size > len(buf) {
		return io.ErrShortWrite
	}

	binary.LittleEndian.PutUint64(buf, u)
	return nil
}

//go:nobounds
func PutUint32(buf []byte, u uint32) error {
	if len(buf)+Uint32Size > len(buf) {
		return io.ErrShortWrite
	}

	binary.LittleEndian.PutUint32(buf, u)
	return nil
}

//go:nobounds
func PutUint16(buf []byte, u uint16) error {
	if len(buf)+Uint16Size > len(buf) {
		return io.ErrShortWrite
	}

	binary.LittleEndian.PutUint16(buf, u)
	return nil
}
