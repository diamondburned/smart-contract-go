package smartcontract

import (
	"./internal/binary"
	"./internal/errors"
	"./internal/utils"
)

type Parameters struct {
	RoundIDX      uint64
	RoundID       [32]byte
	TransactionID [32]byte
	Sender        [32]byte
	Amount        uint64

	buf []byte
}

func Load() (*Parameters, error) {
	// Get the payload length
	payloadLen := _payload_len()

	p := Parameters{
		// Make a byte slice of that length
		buf: make([]byte, payloadLen),
	}

	// Pass in the byte slice, C style
	_payload(&p.buf[0])

	// Read round idx
	u, err := binary.CutUint64(p.buf)
	if err != nil {
		return nil, err
	}

	p.RoundIDX = u

	if err := utils.Cut(p.RoundID[:], p.buf); err != nil {
		return nil, err
	}

	if err := utils.Cut(p.TransactionID[:], p.buf); err != nil {
		return nil, err
	}

	if err := utils.Cut(p.Sender[:], p.buf); err != nil {
		return nil, err
	}

	u, err = binary.CutUint64(p.buf)
	if err != nil {
		return nil, errors.Wrap("Failed to read amount", err)
	}

	p.Amount = u

	return &p, nil
}

func (p *Parameters) Encode() []byte {
	// 8  for uint64   x 2
	// 32 for [32]byte x 3
	buf := make([]byte, 8*2+32*3)

	// Write the idx +8
	if err := binary.PutUint64(buf, p.RoundIDX); err != nil {
		panic(err)
	}

	// Write the id +32
	if err := utils.Copy(buf, p.RoundID[:]); err != nil {
		panic(err) // This shouldn't EVER panic
	}

	// Write the transaction ID +32
	if err := utils.Copy(buf, p.TransactionID[:]); err != nil {
		panic(err)
	}

	// Write the sender +32
	if err := utils.Copy(buf, p.Sender[:]); err != nil {
		panic(err)
	}

	// Write the amount +8
	if err := binary.PutUint64(buf, p.Amount); err != nil {
		panic(err)
	}

	return buf
}
