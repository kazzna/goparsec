package goparsec

type errorMappedParser struct {
	base parserBase
	f    func(ParseError) ParseError
}

func (p *errorMappedParser) parse(t ParseTarget) (ParseResult, ParseTarget, *ParseError) {
	r1, t1, e1 := p.base.parse(t)
	if e1 != nil {
		e2 := p.f(*e1)
		return r1, t1, &e2
	}
	return r1, t1, nil
}
