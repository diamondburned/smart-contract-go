package smartcontract

import "errors"

type HashAlgorithm uint8

const (
	HashBlake2B256 HashAlgorithm = iota
	HashBlake2B512
	HashSHA256
	HashSHA512
)

type SignatureAlgorithm uint8

const (
	SigED25519 SignatureAlgorithm = iota
)

const (
	Blake2B256OutputSize = 32
	Blake2B512OutputSize = 64
	SHA256OutputSize     = 32
	SHA512OutputSize     = 64
)

var (
	ErrVerifyMismatched = errors.New("verification mismatched")
	ErrHashingFailed    = errors.New("hashing failed")
	ErrUnknownAlgorithm = errors.New("unknown algorithm")
)

func (alg SignatureAlgorithm) Verify(pubkey, data, sig []byte) error {
	switch alg {
	case SigED25519:
		if _verify_ed25519(
			&pubkey[0], uint32(len(pubkey)),
			&data[0], uint32(len(data)),
			&sig[0], uint32(len(sig)),
		) != 0 {
			return ErrVerifyMismatched
		}

	default:
		return ErrUnknownAlgorithm
	}

	return nil
}

func (alg HashAlgorithm) Hash(data []byte) ([]byte, error) {
	var (
		out []byte
		err int32
	)

	switch alg {
	case HashBlake2B256:
		out = make([]byte, Blake2B256OutputSize)
		err = _hash_blake2b_256(&data[0], uint32(len(data)), &out[0], uint32(len(out)))

	case HashBlake2B512:
		out = make([]byte, Blake2B512OutputSize)
		err = _hash_blake2b_512(&data[0], uint32(len(data)), &out[0], uint32(len(out)))

	case HashSHA256:
		out = make([]byte, SHA256OutputSize)
		err = _hash_sha256(&data[0], uint32(len(data)), &out[0], uint32(len(out)))

	case HashSHA512:
		out = make([]byte, SHA512OutputSize)
		err = _hash_sha512(&data[0], uint32(len(data)), &out[0], uint32(len(out)))

	default:
		return nil, ErrUnknownAlgorithm
	}

	if err != 0 {
		return out, ErrHashingFailed
	}

	return out, nil
}
