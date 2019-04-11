package tri

import (
	"errors"
)

type Bool int

// Synonyms:
//		Tribool
//		Trilean

const (
	Unknown Bool = iota
	True
	False
)

const (
	strUnknown = "Unknown"
	strTrue    = "True"
	strFalse   = "False"
)

var (
	key_Bool = map[Bool]string{
		Unknown: strUnknown,
		True:    strTrue,
		False:   strFalse,
	}

	val_Bool = map[string]Bool{
		strUnknown: Unknown,
		strTrue:    True,
		strFalse:   False,
	}
)

//---------------
//   Or | F T U |
// -----|-------|
//    F | F T U |
//    T | T T T |
//    U | U T U |
//---------------
func (a Bool) Or(b Bool) (c Bool) {

	switch {
	case (a == True) || (b == True):
		c = True

	case (a == False) && (b == False):
		c = False
	}

	return
}

//---------------
//  And | F T U |
// -----|-------|
//    F | F F F |
//    T | F T U |
//    U | F U U |
//---------------
func (a Bool) And(b Bool) (c Bool) {

	switch {
	case (a == False) || (b == False):
		c = False

	case (a == True) && (b == True):
		c = True
	}

	return
}

//---------------
//  Xor | F T U |
// -----|-------|
//    F | F T U |
//    T | T F U |
//    U | U U U |
//---------------
func (a Bool) Xor(b Bool) (c Bool) {

	switch a {
	case True:
		{
			switch b {
			case True:
				c = False
			case False:
				c = True
			}
		}

	case False:
		{
			switch b {
			case True:
				c = True
			case False:
				c = False
			}
		}
	}

	return
}

func (a Bool) Not() (b Bool) {

	switch a {
	case False:
		b = True
	case True:
		b = False
	}

	return
}

func (a Bool) String() string {
	s, ok := key_Bool[a]
	if !ok {
		s = strUnknown
	}
	return s
}

func ParseBool(s string) (Bool, error) {
	b, ok := val_Bool[s]
	if ok {
		return b, nil
	}
	return Unknown, errors.New("tri.ParseBool: invalid value")
}
