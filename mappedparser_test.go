package goparsec

import (
	"strconv"
	"testing"
)

func TestMap(t *testing.T) {
	p1 := Rune('1')
	parser := p1.Map(func(i interface{}) (interface{}, error) {
		if r, ok := i.(rune); ok {
			n, err := strconv.ParseInt(string([]rune{r}), 10, 0)
			return int(n), err
		}
		panic(assert())
	})
	result, err := parser.Parse("123")
	if err != nil {
		t.Error(err)
		return
	}
	if actual, ok := result.Get().(int); ok {
		AssertIntEqual(t, 1, actual)
		return
	}
	t.Errorf("parse result type shoud be %T but %T found.", int(0), result.Get())
}
