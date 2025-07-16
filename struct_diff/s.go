package struct_diff

type S1 struct {
	A string
}

func (s S1) SubF1() string {
	return s.A
}

func SF1(s S1) string {
	return s.A
}
func SF2(s S1) string {
	return s.SubF1()
}
