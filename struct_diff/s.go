package struct_diff

import (
	"fmt"
	"strconv"
)

type S1 struct {
	B int
}

func (s S1) SubF1(c1 string) string {
	return fmt.Sprintf("%v", s.B) + " :" + c1
}

func SF1(s S1) string {
	return fmt.Sprintf("%v", s.B) + " :" + strconv.Itoa(s.B)
}
func SF2(s S1) string {
	return s.SubF1("test")
}
