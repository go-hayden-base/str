package str

import (
	"testing"
)

func TestIsEmpty(t *testing.T) {
	a := IsEmpty(" ")
	b := IsEmpty(" a b ")
	c := IsEmpty("")
	if !a || b || !c {
		t.Errorf("TestIsEmpty Fail! (a:%t, b:%t, c:%t)", a, b, c)
	}
}

func TestIsInt(t *testing.T) {
	a := IsInt(" 123 ")
	b := IsInt(" +123 ")
	c := IsInt(" -123 ")
	d := IsInt(" 0012 ")
	e := IsInt(" 123.2 ")
	f := IsInt("abc")
	if !a || !b || !c || !d || e || f {
		t.Errorf("TestIsInt Fail! (a:%t, b:%t, c:%t, d:%t, e:%t, f:%t)", a, b, c, d, e, f)
	}
}

func TestIsFloat(t *testing.T) {
	a := IsFloat(" 123.100 ")
	b := IsFloat(" +123.001 ")
	c := IsFloat(" -123.100")
	d := IsFloat(" 0012.00 ")
	e := IsFloat(" 123 ")
	f := IsFloat("abc")
	if !a || !b || !c || !d || e || f {
		t.Errorf("TestIsFloat Fail! (a:%t, b:%t, c:%t, d:%t, e:%t, f:%t)", a, b, c, d, e, f)
	}
}

func TestIsNumber(t *testing.T) {
	a := IsNumber(" 123.100 ")
	b := IsNumber(" +123.001 ")
	c := IsNumber(" -123")
	d := IsNumber(" 0012.00 ")
	e := IsNumber(" #123 ")
	f := IsNumber("abc")
	if !a || !b || !c || !d || e || f {
		t.Errorf("TestIsNumber Fail! (a:%t, b:%t, c:%t, d:%t, e:%t, f:%t)", a, b, c, d, e, f)
	}
}

func TestIsHttpURL(t *testing.T) {
	a := IsHttpURL("http://www.baidu.com")
	b := IsHttpURL("HttP://www.google.com")
	c := IsHttpURL("https://wwww")
	d := IsHttpURL("ftp://abcdef")
	e := IsHttpURL("abcd")
	if !a || !b || !c || d || e {
		t.Errorf("TestIsHttpURL Fail! (a:%t, b:%t, c:%t, d:%t, e:%t)", a, b, c, d, e)
	}
}
