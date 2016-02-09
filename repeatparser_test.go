package goparsec

import (
	"testing"
)

func TestRepeat(t *testing.T) {
	p1 := Rune('a')
	parser := Repeat(p1)
	result, err := parser.Parse("aaabbb")
	if err != nil {
		t.Errorf("Should have no error.")
		return
	}
	expected := []rune{'a', 'a', 'a'}
	if rs, ok := result.Get().([]rune); ok {
		AssertRuneSliceEqual(t, expected, rs)
		return
	}
	AssertTypeEqual(t, expected, result.Get())
}
