package ternary

// Trilean, TriBool

type Trilean int

const (
	Unknown Trilean = iota
	True
	False
)

const (
	strUnknown = "Unknown"
	strTrue    = "True"
	strFalse   = "False"
)

var (
	key_Trilean = map[Trilean]string{
		Unknown: strUnknown,
		True:    strTrue,
		False:   strFalse,
	}

	val_Trilean = map[string]Trilean{
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
func (a Trilean) Or(b Trilean) (c Trilean) {

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
func (a Trilean) And(b Trilean) (c Trilean) {

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
func (a Trilean) Xor(b Trilean) (c Trilean) {

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

func (a Trilean) Not() (b Trilean) {

	switch a {
	case False:
		b = True
	case True:
		b = False
	}

	return
}

func (a Trilean) String() string {

	s, ok := key_Trilean[a]
	if !ok {
		s = strUnknown
	}
	return s
}

func (a *Trilean) Parse(s string) bool {

	v, ok := val_Trilean[s]
	if ok {
		*a = v
	}
	return ok
}
