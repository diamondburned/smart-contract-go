package smartcontract

import (
	"encoding/binary"
)

func ReadBody(buf []byte) (body []byte) {
	// 4 sizeof uint32
	size := binary.LittleEndian.Uint32(buf[:4])
	buf = buf[4:]

	body = buf[:size]
	buf = buf[size:]
	return
}

func WriteBody(dst, src []byte) []byte {
	header := make([]byte, 4)
	binary.LittleEndian.PutUint32(header, uint32(len(src)))

	dst = append(dst, header...)
	dst = append(dst, src...)

	return dst
}
