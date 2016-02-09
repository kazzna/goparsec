package goparsec

type successParserBase int

func Success() StringParser {
	parser := successParserBase(0)
	return &stringParserImpl{&parser}
}

func (p *successParserBase) parse(t ParseTarget) (ParseResult, ParseTarget, *ParseError) {
	return &result{nil}, t, nil
}
