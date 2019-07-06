package main

import sc "github.com/diamondburned/smart-contract-go"

func main() {}

type Transfer struct {
	sc.Contract
}

// todo
func _contract_init(p *sc.Parameters) *Transfer {
	return &Transfer{}
}

func (t *Transfer) _contract_on_money_received(p *sc.Parameters) error {
	sc.SendTransaction(sc.NewTransfer(p.Sender, (p.Amount+1)/2))
	return nil
}
