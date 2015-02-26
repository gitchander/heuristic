package mindshift

import (
	"errors"
	"fmt"
)

func newError(message string) error {
	return errors.New(fmt.Sprint("mindshift: ", message))
}

func newErrorf(format string, a ...interface{}) error {
	return newError(fmt.Sprintf(format, a...))
}
