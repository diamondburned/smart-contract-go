package smartcontract

// This file contains internal functions injected by the wasm vm

func _payload_len() uint32
func _payload(out *byte)

func _log(content *byte, length uint32)
func _result(result *byte, length uint32)

func _send_transaction(tag uint8, payload *byte, payload_len uint32)

func _verify_ed25519(
	pubkey *byte, pubkey_len uint32,
	data *byte, data_len uint32,
	sig *byte, sig_len uint32,
) int32

func _hash_blake2b_256(
	data *byte, data_len uint32,
	out *byte, out_len uint32,
) int32

func _hash_blake2b_512(
	data *byte, data_len uint32,
	out *byte, out_len uint32,
) int32

func _hash_sha256(
	data *byte, data_len uint32,
	out *byte, out_len uint32,
) int32

func _hash_sha512(
	data *byte, data_len uint32,
	out *byte, out_len uint32,
) int32
