package cubspl

import (
	"errors"
	"fmt"
)

var (
	ErrorInputSize       = errors.New("ErrorInputSize")
	ErrorIndexOutOfRange = errors.New("index out of range")
)

func newError(m string) error {
	return errors.New(fmt.Sprintf("cubspl.%s", m))
}

func newErrorf(format string, a ...interface{}) error {
	return newError(fmt.Sprintf(format, a...))
}
