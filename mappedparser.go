package goparsec

type mappedParser struct {
	base parserBase
	f    func(interface{}) (interface{}, error)
}

func (p *mappedParser) parse(t ParseTarget) (ParseResult, ParseTarget, *ParseError) {
	r1, t1, e1 := p.base.parse(t)
	if e1 != nil {
		return r1, t1, e1
	}
	r2, e2 := r1.Map(p.f)
	if e2 != nil {
		err := &ParseError{e2.Error(), t, t1}
		return r2, t1, err
	}
	return r2, t1, nil
}
