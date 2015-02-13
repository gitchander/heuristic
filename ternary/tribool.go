package ternary

// other synonym Trilean

type TriBool int

const (
	Unknown TriBool = iota
	True
	False
)

//---------------------------------------------------------------------------
//   Or | F T U
// -----|------
//    F | F T U
//    T | T T T
//    U | U T U
//---------------------------------------------------------------------------
func (a TriBool) Or(b TriBool) (c TriBool) {

	switch {
	case (a == True) || (b == True):
		c = True

	case (a == False) && (b == False):
		c = False
	}

	return
}

//---------------------------------------------------------------------------
//  And | F T U
// -----|------
//    F | F F F
//    T | F T U
//    U | F U U
//---------------------------------------------------------------------------
func (a TriBool) And(b TriBool) (c TriBool) {

	switch {
	case (a == False) || (b == False):
		c = False

	case (a == True) && (b == True):
		c = True
	}

	return
}

//---------------------------------------------------------------------------
//  Xor | F T U
// -----|------
//    F | F T U
//    T | T F U
//    U | U U U
//---------------------------------------------------------------------------
func (a TriBool) Xor(b TriBool) (c TriBool) {

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

func (a TriBool) Not() (b TriBool) {

	switch a {
	case False:
		b = True
	case True:
		b = False
	}

	return
}

func (a TriBool) String() (s string) {

	switch a {
	case True:
		s = "True"

	case False:
		s = "False"

	default:
		s = "Unknown"
	}

	return
}
