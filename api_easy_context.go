
package fmr

// ParseWithContext returns parse trees for rule <start> at beginning
func (g *Grammar) ParseWithContext(
	context, text string, starts ...string) ([]*Node, error) {
	return g.extractWithContext(
		func(context, text string, starts ...string) ([]*Parse, error) {
			p, err := g.EarleyParseWithContext(context, text, starts...)
			if err != nil {
				return nil, err
			}
			return []*Parse{p}, nil
		}, context, text, starts...)
}

// ParseAnyWithContext returns parse trees for rule <start> at any position
func (g *Grammar) ParseAnyWithContext(
	context, text string, starts ...string) ([]*Node, error) {
	return g.extractWithContext(
		func(context, text string, starts ...string) ([]*Parse, error) {
			p, err := g.EarleyParseAnyWithContext(context, text, starts...)
			if err != nil {
				return nil, err
			}
			return []*Parse{p}, nil
		}, context, text, starts...)
}
