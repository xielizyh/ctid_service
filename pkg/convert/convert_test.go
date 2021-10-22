package convert

import "testing"

type FormatConvTest struct {
	in, out string
}

var UnderscoreToLowerCamelCases = []FormatConvTest{
	{"hello_world", "helloWorld"},
	{"[{\"created_by\":16}]", "[{\"createdBy\":16}]"},
}

func TestUnderscoreToLowerCamelCase(t *testing.T) {
	for _, s := range UnderscoreToLowerCamelCases {
		exp := StrTo(s.in).UnderscoreToLowerCamelCase()
		if s.out != exp {
			t.Errorf("UnderscoreToLowerCamelCase of %s expects %s, but got %s", s.in, s.out, exp)
		}
	}
}
