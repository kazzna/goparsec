package goparsec

import "fmt"

// parser type method
type parserBase interface {
	// parse input
	parse(ParseTarget) (ParseResult, ParseTarget, *ParseError)
}

//  Parser Interface for string.
type StringParser interface {
	parser() parserBase
	Parse(string) (ParseResult, error)
	Then(StringParser, func(interface{}, interface{}) (interface{}, error)) StringParser
	Map(f func(interface{}) (interface{}, error)) StringParser
	ErrorMap(func(ParseError) ParseError) StringParser
}

type stringParsable interface {
	Length() int
	Head() rune
	Tail() stringParsable
	fmt.Stringer
}

type stringParserImpl struct {
	base parserBase
}

func (p *stringParserImpl) parser() parserBase {
	return p.base
}

func (p *stringParserImpl) Parse(s string) (ParseResult, error) {
	result, _, err := p.parser().parse(RuneListOf(s))
	if err != nil {
		return result, err
	}
	return result, nil
}

func (p *stringParserImpl) Then(p2 StringParser, merge func(interface{}, interface{}) (interface{}, error)) StringParser {
	return &stringParserImpl{&thenParser{p.parser(), p2.parser(), merge}}
}

func (p *stringParserImpl) Map(f func(interface{}) (interface{}, error)) StringParser {
	return &stringParserImpl{&mappedParser{p.parser(), f}}
}

func (p *stringParserImpl) ErrorMap(f func(ParseError) ParseError) StringParser {
	return &stringParserImpl{&errorMappedParser{p.parser(), f}}
}
