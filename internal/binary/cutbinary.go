package binary

import "io"

//go:nobounds
func CutUint64(buf []byte) (uint64, error) {
	u, err := Uint64(buf)
	if err != nil {
		return u, io.ErrUnexpectedEOF
	}

	buf = buf[Uint64Size:]

	return u, nil
}

//go:nobounds
func CutUint32(buf []byte) (uint32, error) {
	u, err := Uint32(buf)
	if err != nil {
		return u, io.ErrUnexpectedEOF
	}

	buf = buf[Uint32Size:]

	return u, nil
}

//go:nobounds
func CutUint16(buf []byte) (uint16, error) {
	u, err := Uint16(buf)
	if err != nil {
		return u, io.ErrUnexpectedEOF
	}

	buf = buf[Uint16Size:]

	return u, nil
}
