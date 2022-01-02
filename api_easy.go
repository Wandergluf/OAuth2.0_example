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

// ParseAny returns parse trees for rule <start> at any position
func (g *Grammar) ParseAny(text string, starts ...string) ([]*Node, error) {
	return g.extract(
		func(text string, starts ...string) ([]*Parse, error) {
			p, err := g.EarleyParseAny(text, starts...)
			if err != nil {
				return nil, err
			}
			return []*Parse{p}, nil
		}, text, starts...)
}

// ExtractMaxAll extracts all parse trees in text for rule <start>
func (g *Grammar) ExtractMaxAll(
	text string, starts ...string) ([]*Node, error) {
	return g.extract(g.EarleyParseMaxAll, text, starts...)
}

// ExtractAll extra