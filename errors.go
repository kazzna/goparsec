package goparsec

import (
	"fmt"
)

// error should never happen without bugs.
func assert() error {
	return fmt.Errorf("-*- Should never comes here. -*-")
}

type ParseTarget interface {
	Length() int
	Head() interface{}
	Tail() ParseTarget
	fmt.Stringer
}

type ParseError struct {
	Message string
	Input ParseTarget
	Next ParseTarget
}

func (p *ParseError) Error() string {
	return p.Message
}
