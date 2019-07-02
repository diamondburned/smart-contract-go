package errors

import "errors"

var _ error = (*_err)(nil)

type _err struct {
	msg string
	err error
}

func Wrap(s string, err error) error {
	return errors.New(s + ": " + err.Error())
}

func (e _err) Unwrap() error {
	return e.err
}

func (e _err) Error() string {
	return e.msg + ": " + e.err.Error()
}
