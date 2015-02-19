package ternary

import "testing"

const anyOther privTriBool = 100 // -> Unknown

type unaryOperatorSample struct {
	a, b privTriBool // b= a.operator()
}

type binaryOperatorSample struct {
	a, b, c privTriBool // c= a.operator(b)
}

var not_Samples = []unaryOperatorSample{

	{privTrue, privFalse},
	{privFalse, privTrue},
	{privUnknown, privUnknown},
	{anyOther, privUnknown},
}

var and_Samples = []binaryOperatorSample{

	{privTrue, privTrue, privTrue},
	{privFalse, privTrue, privFalse},
	{privUnknown, privTrue, privUnknown},
	{anyOther, privTrue, privUnknown},

	{privTrue, privFalse, privFalse},
	{privFalse, privFalse, privFalse},
	{privUnknown, privFalse, privFalse},
	{anyOther, privFalse, privFalse},

	{privTrue, privUnknown, privUnknown},
	{privFalse, privUnknown, privFalse},
	{privUnknown, privUnknown, privUnknown},
	{anyOther, privUnknown, privUnknown},

	{privTrue, anyOther, privUnknown},
	{privFalse, anyOther, privFalse},
	{privUnknown, anyOther, privUnknown},
	{anyOther, anyOther, privUnknown},
}

var or_Samples = []binaryOperatorSample{

	{privTrue, privTrue, privTrue},
	{privFalse, privTrue, privTrue},
	{privUnknown, privTrue, privTrue},
	{anyOther, privTrue, privTrue},

	{privTrue, privFalse, privTrue},
	{privFalse, privFalse, privFalse},
	{privUnknown, privFalse, privUnknown},
	{anyOther, privFalse, privUnknown},

	{privTrue, privUnknown, privTrue},
	{privFalse, privUnknown, privUnknown},
	{privUnknown, privUnknown, privUnknown},
	{anyOther, privUnknown, privUnknown},

	{privTrue, anyOther, privTrue},
	{privFalse, anyOther, privUnknown},
	{privUnknown, anyOther, privUnknown},
	{anyOther, anyOther, privUnknown},
}

var xor_Samples = []binaryOperatorSample{

	{privTrue, privTrue, privFalse},
	{privFalse, privTrue, privTrue},
	{privUnknown, privTrue, privUnknown},
	{anyOther, privTrue, privUnknown},

	{privTrue, privFalse, privTrue},
	{privFalse, privFalse, privFalse},
	{privUnknown, privFalse, privUnknown},
	{anyOther, privFalse, privUnknown},

	{privTrue, privUnknown, privUnknown},
	{privFalse, privUnknown, privUnknown},
	{privUnknown, privUnknown, privUnknown},
	{anyOther, privUnknown, privUnknown},

	{privTrue, anyOther, privUnknown},
	{privFalse, anyOther, privUnknown},
	{privUnknown, anyOther, privUnknown},
	{anyOther, anyOther, privUnknown},
}

func TestTriBool(t *testing.T) {

	var a, b, c TriBool

	for _, s := range not_Samples {

		a.v = s.a
		b.v = s.b

		if !a.Not().Equal(b) {
			t.Errorf("not %s", s.a)
		}
	}

	for _, s := range and_Samples {

		a.v = s.a
		b.v = s.b
		c.v = s.c

		if !a.And(b).Equal(c) {
			t.Errorf("%s and %s", s.a, s.b)
		}
	}

	for _, s := range or_Samples {

		a.v = s.a
		b.v = s.b
		c.v = s.c

		if !a.Or(b).Equal(c) {
			t.Errorf("%s or %s", s.a, s.b)
		}
	}

	for _, s := range xor_Samples {

		a.v = s.a
		b.v = s.b
		c.v = s.c

		if !a.Xor(b).Equal(c) {
			t.Errorf("%s xor %s", s.a, s.b)
		}
	}
}
