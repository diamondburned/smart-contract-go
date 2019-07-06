package encode

import (
	"encoding/binary"
	"io"

	"../utils"
)

// CutBody pops the body out, slicing the buf.
//go:nobounds
func CutBody(buf []byte) (body []byte, err error) {
	if len(buf) < 4 {
		return nil, io.ErrUnexpectedEOF
	}

	// 4 sizeof uint32
	size := int(binary.LittleEndian.Uint32(buf[:4]))
	buf = buf[4:]

	if len(buf) < size {
		return nil, io.ErrUnexpectedEOF
	}

	body = buf[:size]
	buf = buf[size:]

	return
}

// WriteBody writes max_len(dst) bytes for the length header and the source
// bytes body. An `io.ErrShortWrite' is returned if `dst' is not long enough.
func WriteBody(dst, src []byte) error {
	header := make([]byte, 4)
	binary.LittleEndian.PutUint32(header, uint32(len(src)))

	if err := utils.Copy(dst, header); err != nil {
		return io.ErrShortWrite
	}

	if err := utils.Copy(dst, src); err != nil {
		return io.ErrShortWrite
	}

	return nil
}

// AppendBody appends the header and the body straight into the bytes array then
// return it. It is encouraged for one to pre-allocate a byte slice to zeros
// then use WriteBody instead of using AppendBody for optimization. This returns
// no errors, as it only appends the bytes, thus does not care if the slice has
// enough space.
func AppendBody(dst, src []byte) []byte {
	header := make([]byte, 4)
	binary.LittleEndian.PutUint32(header, uint32(len(src)))

	dst = append(dst, header...)
	dst = append(dst, src...)

	return dst
}
