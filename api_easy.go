package fmr

// Parse returns parse trees for rule <start> at beginning
func (g *Grammar) Parse(text string, starts ...string) ([]*Node, error) {
	return g.e