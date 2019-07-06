package main

import (
	sc "github.com/diamondburned/smart-contract-go"
	"github.com/diamondburned/smart-contract-go/internal/utils"
)

func main() {}

type Token struct {
	// An array is used as tinygo has very limited hashmap support
	Balances []Balance
}

type Balance struct {
	Sender [32]byte
	Amount uint64
}

func _contract_init(p *sc.Parameters) *Token {
	return &Token{}
}

func (t *Token) _contract_balance(p *sc.Parameters) error {
	var amount uint64

	if b := t.get(p.Sender); b != nil {
		amount = b.Amount
	}

	utils.Debug(p.Sender, amount)
	return nil
}

func (t *Token) _contract_transfer(p *sc.Parameters) error {
	return nil
}

// internal getter and setter to replace hashmaps

func (t *Token) get(sender [32]byte) *Balance {
	for _, b := range t.Balances {
		if b.Sender == sender {
			return &b
		}
	}

	return nil
}

func (t *Token) set(sender [32]byte, amount uint64) {
	for i, b := range t.Balances {
		if b.Sender == sender {
			t.Balances[i].Amount = amount
			return
		}
	}

	t.Balances = append(t.Balances, Balance{sender, amount})
}

func println(s)
