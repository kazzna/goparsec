package goparsec

type ParseResult interface {
	// Convert result value by parameter function.
	Map(func(interface{}) (interface{}, error)) (ParseResult, error)
	// Get the result value.
	Get() interface{}
}

type result struct {
	value interface{}
}

func (p *result) Map(f func(interface{}) (interface{}, error)) (ParseResult, error) {
	r, err := f(p.value)
	if err != nil {
		return nil, err
	}
	return &result{r}, nil
}

func (p *result) Get() interface{} {
	return p.value
}
