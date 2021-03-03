package tri

import (
	"errors"
)

// Synonyms:
// Tribool
// Trilean

type Tri int

const (
	False   Tri = -1
	Unknown Tri = 0
	True    Tri = 1
)

const (
	strUnknown = "Unknown"
	strTrue    = "True"
	strFalse   = "False"
)

var (
	keyTri = map[Tri]string{
		Unknown: strUnknown,
		True:    strTrue,
		False:   strFalse,
	}

	valTri = map[string]Tri{
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
func (a Tri) Or(b Tri) (c Tri) {

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
func (a Tri) And(b Tri) (c Tri) {

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
func (a Tri) Xor(b Tri) (c Tri) {

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

func (a Tri) Not() (b Tri) {

	switch a {
	case False:
		b = True
	case True:
		b = False
	}

	return
}

func (a Tri) String() string {
	s, ok := keyTri[a]
	if !ok {
		s = strUnknown
	}
	return s
}

func ParseTri(s string) (Tri, error) {
	b, ok := valTri[s]
	if ok {
		return b, nil
	}
	return Unknown, errors.New("tri.ParseTri: invalid value")
}
