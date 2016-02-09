package goparsec

type thenParser struct {
	before parserBase
	after  parserBase
	merge  func(interface{}, interface{}) (interface{}, error)
}

func (p *thenParser) parse(t ParseTarget) (ParseResult, ParseTarget, *ParseError) {
	pr1, t1, e1 := p.before.parse(t)
	if e1 != nil {
		return nil, t1, e1
	}
	pr2, t2, e2 := p.after.parse(t1)
	if e2 != nil {
		return nil, t2, e2
	}
	r, err := p.merge(pr1.Get(), pr2.Get())
	if err != nil {
		e3 := &ParseError{err.Error(), t, t2}
		return &result{r}, t2, e3
	}
	return &result{r}, t2, nil
}
