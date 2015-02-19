package ternary

type privTriBool int

const (
	privUnknown privTriBool = iota // default value
	privTrue
	privFalse
)

const (
	strUnknown = "Unknown"
	strTrue    = "True"
	strFalse   = "False"
)

var (
	key_TriBool = map[privTriBool]string{
		privUnknown: strUnknown,
		privTrue:    strTrue,
		privFalse:   strFalse,
	}

	val_TriBool = map[string]privTriBool{
		strUnknown: privUnknown,
		strTrue:    privTrue,
		strFalse:   privFalse,
	}
)

// other synonym Trilean
type TriBool struct {
	v privTriBool // hide operators '==' and '='
}

func (a *TriBool) SetTrue() *TriBool {
	a.v = privTrue
	return a
}

func (a *TriBool) SetFalse() *TriBool {
	a.v = privFalse
	return a
}

func (a *TriBool) SetUnknown() *TriBool {
	a.v = privUnknown
	return a
}

func (a TriBool) IsTrue() bool {
	return (a.v == privTrue)
}

func (a TriBool) IsFalse() bool {
	return (a.v == privFalse)
}

func (a TriBool) IsUnknown() bool {
	return (a.v == privUnknown)
}

func (a TriBool) Equal(b TriBool) bool {

	return (a.v == b.v)
}

//---------------
//   Or | F T U |
// -----|-------|
//    F | F T U |
//    T | T T T |
//    U | U T U |
//---------------
func (a TriBool) Or(b TriBool) (c TriBool) {

	switch {
	case (a.v == privTrue) || (b.v == privTrue):
		c.v = privTrue

	case (a.v == privFalse) && (b.v == privFalse):
		c.v = privFalse
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
func (a TriBool) And(b TriBool) (c TriBool) {

	switch {
	case (a.v == privFalse) || (b.v == privFalse):
		c.v = privFalse

	case (a.v == privTrue) && (b.v == privTrue):
		c.v = privTrue
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
func (a TriBool) Xor(b TriBool) (c TriBool) {

	switch a.v {
	case privTrue:
		{
			switch b.v {
			case privTrue:
				c.v = privFalse
			case privFalse:
				c.v = privTrue
			}
		}

	case privFalse:
		{
			switch b.v {
			case privTrue:
				c.v = privTrue
			case privFalse:
				c.v = privFalse
			}
		}
	}

	return
}

func (a TriBool) Not() (b TriBool) {

	switch a.v {
	case privFalse:
		b.v = privTrue
	case privTrue:
		b.v = privFalse
	}

	return
}

func (a TriBool) String() string {

	s, ok := key_TriBool[a.v]
	if !ok {
		s = strUnknown
	}
	return s
}

func (a *TriBool) Parse(s string) bool {

	v, ok := val_TriBool[s]
	if ok {
		a.v = v
	}
	return ok
}
