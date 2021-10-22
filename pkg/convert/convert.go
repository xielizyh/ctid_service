package convert

import (
	"strconv"
	"strings"
	"unicode"
)

type StrTo string

func (s StrTo) String() string {
	return string(s)
}

func (s StrTo) Int() (int, error) {
	v, err := strconv.Atoi(s.String())
	return v, err
}

func (s StrTo) MustInt() int {
	v, _ := s.Int()
	return v
}

func (s StrTo) UInt32() (uint32, error) {
	v, err := strconv.Atoi(s.String())
	return uint32(v), err
}

func (s StrTo) MustUInt32() uint32 {
	v, _ := s.UInt32()
	return v
}

func (s StrTo) UnderscoreToUpperCamelCase() string {
	v := strings.Replace(s.String(), "_", " ", -1)
	v = strings.Title(v)
	return strings.Replace(v, " ", "", -1)
}

func (s StrTo) UnderscoreToLowerCamelCase() string {
	v := s.UnderscoreToUpperCamelCase()
	return string(unicode.ToLower(rune(v[0]))) + v[1:]
}
