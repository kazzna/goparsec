package goparsec

import (
	"fmt"
	"reflect"
	"testing"
)

func makeError(t *testing.T, expected, actual string) {
	t.Errorf("%s expected but actual: %s.", expected, actual)
}

func AssertTypeEqual(t *testing.T, expected, actual interface{}) {
	if reflect.TypeOf(actual) != reflect.TypeOf(expected) {
		f := "%t"
		makeError(t, fmt.Sprintf(f, expected), fmt.Sprintf(f, actual))
	}
}

func AssertStringEqual(t *testing.T, expected, actual string) {
	if actual != expected {
		f := "\"%s\""
		makeError(t, fmt.Sprintf(f, expected), fmt.Sprintf(f, actual))
	}
}

func AssertIntEqual(t *testing.T, expected, actual int) {
	if actual != expected {
		f := "%d"
		makeError(t, fmt.Sprintf(f, expected), fmt.Sprintf(f, actual))
	}
}

func AssertRuneEqual(t *testing.T, expected, actual rune) {
	if actual != expected {
		f := "'%c'"
		makeError(t, fmt.Sprintf(f, expected), fmt.Sprintf(f, actual))
	}
}

func AssertRuneSliceEqual(t *testing.T, expected, actual []rune) {
	la := len(actual)
	le := len(expected)
	if la != le {
		f := "'Slice length %d'"
		makeError(t, fmt.Sprintf(f, le), fmt.Sprintf(f, la))
	}
	for i := 0; i < le; i++ {
		if actual[i] != expected[i] {
			t.Errorf("At i: %s expected but actual: %s.", i, expected[i], actual[i])
		}
	}
}

func runeToString(i interface{}) string {
	if r, ok := i.(rune); ok {
		return string([]rune{r})
	}
	panic(fmt.Errorf("ParseError"))
}

func stringToString(i interface{}) string {
	if r, ok := i.(string); ok {
		return r
	}
	panic(fmt.Errorf("ParseError"))
}

func rune2ToString(i, j interface{}) (interface{}, error) {
	if r1, ok := i.(rune); ok {
		if r2, ok := j.(rune); ok {
			return string([]rune{r1, r2}), nil
		}
		return nil, fmt.Errorf("j is not rune")
	}
	return nil, fmt.Errorf("i is not rune")
}

func TestRune(t *testing.T) {
	expected := 'あ'
	p := Rune(expected)
	result, err := p.Parse("あいう")
	if err != nil {
		fmt.Printf("エラー内容: %s", err.Error())
		t.Error(err)
		return
	}
	if actual, ok := result.Get().(rune); ok {
		AssertRuneEqual(t, expected, actual)
		return
	}
	t.Errorf("parse result is not string")
}

func TestRuneThenRune(t *testing.T) {
	p1 := Rune('日')
	p2 := Rune('本')
	p := p1.Then(p2, rune2ToString)
	result, err := p.Parse("日本語")
	if err != nil {
		t.Error(err)
		return
	}
	if actual, ok := result.Get().(string); ok {
		AssertStringEqual(t, "日本", actual)
		return
	}
	t.Errorf("parse result is not string")
}

func TestString(t *testing.T) {
	expected := "日本語"
	parser := String(expected)
	result, err := parser.Parse("日本語学習")
	if err != nil {
		t.Error(err)
		return
	}
	if actual, ok := result.Get().(string); ok {
		AssertStringEqual(t, expected, actual)
		return
	}
	t.Errorf("parse result is not string")
}

func TestStringErrorMessage(t *testing.T) {
	expected := "\"日本語\" expected but \"日本人街\" found."
	parser := String("日本語")
	_, err := parser.Parse("日本人街")
	if err != nil {
		AssertStringEqual(t, expected, err.Error())
		return
	}
	t.Errorf("Error should occur.")
}

func TestEOS(t *testing.T) {
	expected := "End of string expected but aaa found."
	parser := EOS()
	_, err := parser.Parse("aaa")
	if err == nil {
		t.Errorf("error expected but not returned.")
		return
	}
	AssertStringEqual(t, expected, err.Error())
}
