package smartcontract

type Contract struct {
	Payload []byte
	Code    []byte
}

func ReadFromContract(buf []byte) *Contract {
	return &Contract{
		Payload: ReadBody(buf),
		Code:    buf,
	}
}

func (c *Contract) WriteTo(buf []byte) {
	buf = WriteBody(buf, c.Payload)
	buf = append(buf, c.Code...)
}
