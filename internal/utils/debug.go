package utils

import (
	"encoding/base64"
	"strconv"
)

func Debug(i ...interface{}) {
	for _, i := range i {
		switch i := i.(type) {
		case [32]byte:
			buf := make([]byte, 32)
			base64.StdEncoding.Encode(buf, i[:])
			print(string(buf))
		case [64]byte:
			buf := make([]byte, 64)
			base64.StdEncoding.Encode(buf, i[:])
			print(string(buf))
		case []byte:
			print(string(i))
		case string:
			print(i)
		case int:
			print(strconv.Itoa(i))
		case int64:
			print(strconv.FormatInt(i, 10))
		case int32:
			print(strconv.FormatInt(int64(i), 10))
		case int16:
			print(strconv.FormatInt(int64(i), 10))
		case int8:
			print(strconv.FormatInt(int64(i), 10))
		case uint:
			print(strconv.FormatUint(uint64(i), 10))
		case uint64:
			print(strconv.FormatUint(i, 10))
		case uint32:
			print(strconv.FormatUint(uint64(i), 10))
		case uint16:
			print(strconv.FormatUint(uint64(i), 10))
		case uint8:
			print(strconv.FormatUint(uint64(i), 10))
		case float32:
			print(strconv.FormatFloat(float64(i), 'e', -1, 32))
		case float64:
			print(strconv.FormatFloat(float64(i), 'e', -1, 64))
		case error:
			print(i.Error())
		default:
			print(i)
		}

		print(" ")
	}
}

func Debugln(i ...interface{}) {
	Debug(append(i, "\n")...)
}
