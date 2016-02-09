package goparsec

type repeatParserBase struct {
	base parserBase
}

func (p *repeatParserBase) parse(t ParseTarget) (ParseResult, ParseTarget, *ParseError) {
	res := make([]interface{}, 0)
	t1 := t
	for cont := true; cont; {
		r1, t2, e1 := p.base.parse(t1)
		if e1 != nil {
			cont = false
		} else {
			t1 = t2
			res = append(res, r1.Get())
		}
	}
	return &result{res}, t1, nil
}

func Repeat(parser StringParser) StringParser {
	return &stringParserImpl{&repeatParserBase{parser.parser()}}
}
