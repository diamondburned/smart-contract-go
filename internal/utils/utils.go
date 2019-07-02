package utils

import (
	"io"
)

// Copy tries to fully copy() and return an error if it cannot copy fully.
func Copy(dst, src []byte) error {
	if copy(dst, src) != len(src) {
		return io.ErrShortWrite
	}

	return nil
}

// Cut copies then trims `src', effectively "cutting."
func Cut(dst, src []byte) error {
	if err := Copy(dst, src); err != nil {
		return err
	}

	// Copy would fail if dst is smaller than src, thus it is safe to just slice
	// src
	src = src[len(dst):]

	return nil
}
