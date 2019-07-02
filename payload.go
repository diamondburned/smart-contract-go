package smartcontract

import "encoding/binary"

type Parameters struct {
	RoundIDX      uint64
	RoundID       [32]byte
	TransactionID [32]byte
	Sender        [32]byte
	Amount        uint64
}

func Load() *Parameters {
	// Get the payload length
	payloadLen := _payload_len()

	// Make a byte slice of that length
	payload := make([]byte, payloadLen)

	// Pass in the byte slice, C style
	_payload(&payload[0])

	p := Parameters{}

	// Read round idx
	p.RoundIDX = binary.LittleEndian.Uint64(payload)
	payload = payload[8:]

	copy(p.RoundID[:], payload)
	payload = payload[32:]

	copy(p.TransactionID[:], payload)
	payload = payload[32:]

	copy(p.Sender[:], payload)
	payload = payload[32:]

	p.Amount = binary.LittleEndian.Uint64(payload)
	payload = payload[8:]

	return &p
}
