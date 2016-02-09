package goparsec

import (
	"fmt"
)

type runeParserBase struct {
	value rune
}

func (p *runeParserBase) parse(t ParseTarget) (ParseResult, ParseTarget, *ParseError) {
	if head, ok := t.Head().(rune); ok {
		tail := t.Tail()
		if head != p.value {
			fmt.Printf("15行目: %s\n", t.String())
			msg := fmt.Sprintf("'%c' expected but '%c' found.\n", p.value, []rune(t.String())[0])
			err := &ParseError{msg, t, t}
			return nil, tail, err
		}
		return &result{p.value}, tail, nil
	}
	msg := fmt.Sprintf("ParseTarget(%s)'s Head is not rune.\n", t.String())
	err := &ParseError{msg, t, t}
	fmt.Printf("24行目: %s\n", msg)
	return nil, t, err
}

func Rune(r rune) StringParser {
	return &stringParserImpl{&runeParserBase{r}}
}

type eosParserBase int

func (p *eosParserBase) parse(t ParseTarget) (ParseResult, ParseTarget, *ParseError) {
	if t.Length() != 0 {
		msg := fmt.Sprintf("End of string expected but %s found.", t.String())
		err := &ParseError{msg, t, t}
		return nil, t, err
	}
	return &result{nil}, t, nil
}

func EOS() StringParser {
	r := eosParserBase(0)
	return &stringParserImpl{&r}
}

type stringParserBase struct {
	value string
}

func (p *stringParserBase) parse(t ParseTarget) (ParseResult, ParseTarget, *ParseError) {
	parser := Success().Map(func(_ interface{}) (interface{}, error) {
		return make([]rune, 0), nil
	})
	for _, r := range p.value {
		parser = parser.Then(Rune(r), func(i, j interface{}) (interface{}, error) {
			if rs, ok := i.([]rune); ok {
				if r, ok := j.(rune); ok {
					return append(rs, r), nil
				}
			}
			panic(assert())
		})
	}

	parser = parser.Map(func(i interface{}) (interface{}, error) {
		if rs, ok := i.([]rune); ok {
			return string(rs), nil
		}
		panic(assert())
	})

	r1, t1, e1 := parser.parser().parse(t)
	if e1 != nil {
		msg := fmt.Sprintf("\"%s\" expected but \"%s\" found.", p.value, t.String())
		err := &ParseError{msg, t, t}
		return r1, t, err
	}
	return r1, t1, nil
}

func String(s string) StringParser {
	return &stringParserImpl{&stringParserBase{s}}
}
