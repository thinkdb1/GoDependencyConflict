package struct_diff

import "strconv"

type S1 struct {
	A string
	B int
}

func (s S1) SubF1(c1 string) string {
	return s.A + " :" + c1
}

func SF1(s S1) string {
	return s.A + " :" + strconv.Itoa(s.B)
}
func SF2(s S1) string {
	return s.SubF1("test")
}
