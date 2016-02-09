package goparsec

type RuneList struct {
	values []rune
}

func RuneListOf(s string) *RuneList {
	return &RuneList{[]rune(s)}
}

func (rs *RuneList) Length() int {
	return len(rs.values)
}

func (rs *RuneList) Head() interface{} {
	return rs.values[0]
}

func (rs *RuneList) Tail() ParseTarget {
	return &RuneList{rs.values[1:]}
}

func (rs *RuneList) String() string {
	return string(rs.values)
}
