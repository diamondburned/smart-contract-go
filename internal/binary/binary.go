// Package binary implements bounds-safe read functions for Little Endian
// encoded bytes int (unsigned) integer types. All functions would reslice the
// buffer so that it doesn't have the same bytes after running.
package binary

import (
	"encoding/binary"
	"io"
)

const (
	Uint64Size = 8
	Uint32Size = 4
	Uint16Size = 2
	Uint8Size  = 1
)

//go:nobounds
func Uint64(buf []byte) (uint64, error) {
	if len(buf) < Uint64Size {
		return 0, io.ErrUnexpectedEOF
	}

	return binary.LittleEndian.Uint64(buf), nil
}

//go:nobounds
func Uint32(buf []byte) (uint32, error) {
	if len(buf) < Uint32Size {
		return 0, io.ErrUnexpectedEOF
	}

	return binary.LittleEndian.Uint32(buf), nil
}

//go:nobounds
func Uint16(buf []byte) (uint16, error) {
	if len(buf) < Uint16Size {
		return 0, io.ErrUnexpectedEOF
	}

	return binary.LittleEndian.Uint16(buf), nil
}
