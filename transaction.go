package smartcontract

import (
	"io"

	"./internal/binary"
	"./internal/encode"
	"./internal/errors"
	"./internal/utils"
)

type TransactionTag uint8

const (
	TagNop TransactionTag = iota
	TagTransfer
	TagContract
	TagStake
)

type Transaction interface {
	Encode() []byte
	Tag() TransactionTag
}

func SendTransaction(t Transaction) {
	b := t.Encode()
	_send_transaction(byte(t.Tag()), &b[0], uint32(len(b)))
}

type Transfer struct {
	Destination [32]byte
	Amount      uint64

	FuncName   string
	FuncParams string
}

func NewTransfer(destination [32]byte, amount uint64) *Transfer {
	return &Transfer{destination, amount, "", ""}
}

func ReadFromTransfer(buf []byte) (*Transfer, error) {
	t := Transfer{}

	if err := utils.Cut(t.Destination[:], buf); err != nil {
		return nil, errors.Wrap("Failed to write the destination", err)
	}

	a, err := binary.CutUint64(buf)
	if err != nil {
		return nil, errors.Wrap("Failed to read the amount", err)
	}

	t.Amount = a

	if len(buf) > 0 {
		b, err := encode.CutBody(buf)
		if err != nil {
			return nil, errors.Wrap("Failed to read function name", err)
		}

		t.FuncName = string(b)
	}

	if len(buf) > 0 {
		b, err := encode.CutBody(buf)
		if err != nil {
			return nil, errors.Wrap("Failed to read function params", err)
		}

		t.FuncParams = string(b)
	}

	return &t, nil
}

func (t *Transfer) SetFunction(name, params string) {
	t.FuncName = name
	t.FuncParams = params
}

// Methods to satisfy the Transaction interface
var _ Transaction = (*Transfer)(nil)

func (t *Transfer) Tag() TransactionTag {
	return TagTransfer
}

func (t *Transfer) Encode() []byte {
	buf := make([]byte,
		len(t.Destination)+8*2+len(t.FuncName)+len(t.FuncParams),
	)

	// Write the destination
	if err := utils.Copy(buf, t.Destination[:]); err != nil {
		panic(errors.Wrap("Failed to write the destination", io.ErrShortWrite))
	}

	// Write the amount
	if err := binary.PutUint64(buf, t.Amount); err != nil {
		panic(errors.Wrap("Failed to write the amount", err))
	}

	// Write a dummy gas amount
	if err := binary.PutUint64(buf, 0); err != nil {
		panic(errors.Wrap("Failed to write the gas amount", err))
	}

	if len(t.FuncName) > 0 && len(t.FuncParams) > 0 {
		if err := utils.Copy(buf, []byte(t.FuncName)); err != nil {
			panic(errors.Wrap("Failed to write function name", io.ErrShortWrite))
		}

		if err := utils.Copy(buf, []byte(t.FuncParams)); err != nil {
			panic(errors.Wrap("Failed to write function params", io.ErrShortWrite))
		}
	}

	return buf
}
