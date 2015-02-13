package hexm

import (
	"errors"
	"fmt"
)

func newError(m string) error {
	return errors.New(fmt.Sprintf("hexm.%s", m))
}

func newErrorCoord(m string) error {
	return newError(fmt.Sprintf("Coord.%s", m))
}

var (
	ErrorNeighborDirInvalid     = newError("Neighbor dir invalid")
	ErrorCoordNegativeParameter = newErrorCoord("Parameter can not be negative")
	ErrorCoordOneZeroParameter  = newErrorCoord("At least one parameter must be zero")
	ErrorSizeOutOfRange         = newError("Size out of range")
	ErrorSizeZeroParameter      = newError("Parameters can not be negative or zero")
	ErrorIndexIsNotValit        = newError("Index is not valid")
	ErrorIteratorIndexOut       = newError("Iterator index out")
)