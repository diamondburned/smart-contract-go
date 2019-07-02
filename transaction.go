package smartcontract

import "encoding/binary"

type TransactionTag uint8

const (
	TagNop TransactionTag = iota
	TagTransfer
	TagContract
	TagStake
)

type Transaction interface {
	Write() []byte
	Tag() TransactionTag
}

func SendTransaction(t Transaction) {
	buf := t.Write()
	_send_transaction(byte(t.Tag()), &buf[0], uint32(len(buf)))
}

type Transfer struct {
	Destination [32]byte
	Amount      uint64

	FuncName   string
	FuncParams string
}

func ReadFromTransfer(buf []byte) *Transfer {
	t := Transfer{}

	copy(t.Destination[:], buf)
	buf = buf[32:]

	t.Amount = binary.LittleEndian.Uint64(buf)
	buf = buf[4:]

	if len(buf) > 0 {
		t.FuncName = string(ReadBody(buf))
	}

	if len(buf) > 0 {
		t.FuncParams = string(ReadBody(buf))
	}

	return &t
}

var _ Transaction = (*Transfer)(nil)

func (t *Transfer) Tag() TransactionTag {
	return TagTransfer
}

func (t *Transfer) Write() (buf []byte) {
	buf = make([]byte, 0,
		len(t.Destination)+8+len(t.FuncName)+len(t.FuncParams),
	)

	// Write the destination
	buf = append(buf, t.Destination[:]...)

	// Write the amount
	binary.LittleEndian.PutUint64(buf, t.Amount)

	if len(t.FuncName) > 0 && len(t.FuncParams) > 0 {
		buf = append(buf, []byte(t.FuncName)...)
		buf = append(buf, []byte(t.FuncParams)...)
	}

	return
}
