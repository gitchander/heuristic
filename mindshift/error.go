package mindshift

import (
	"errors"
	"fmt"
)

func newError(message string) error {
	return errors.New(fmt.Sprint("mindshift: ", message))
}
