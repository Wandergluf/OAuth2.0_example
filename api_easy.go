package fmr

// Parse returns parse trees for rule <start> at beginning
func (g *Grammar) Parse(text string, starts ...string) ([]*Node, error) {
	return g.extract(func(text string, starts ...string) ([]*Parse, error) {
		p, err := g.EarleyParse(text, starts...)
		if err != nil {
			return nil, err
		}
		return []*Parse{p}, nil
	}, text, starts...)
}

// ParseAny returns parse trees fo