package smartcontract

import "./internal/encode"

type Contract struct {
	Payload []byte
	Code    []byte
}

func ReadFromContract(buf []byte) (*Contract, error) {
	b, err := encode.CutBody(buf)
	if err != nil {
		return nil, err
	}

	return &Contract{
		Payload: b,
		Code:    buf,
	}, nil
}

func (c *Contract) WriteTo(buf []byte) {
	buf = encode.AppendBody(buf, c.Payload)
	buf = append(buf, c.Code...)
}
