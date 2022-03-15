package fmr

import "fmt"

func (p *parser) special() (*Term, error) {
	if err := p.eat('('); err != nil {
		return nil, err
	}
	p.ws()
	name, err := p.text()
	if err != nil {
		return nil, err
	}
	p.ws()
	switch name {
	case "any":
		return p.any()
	case "list":
		return p.list()
	default:
		return nil, fmt.Errorf(
			"%s: special rule:(%s) not supported", p.posInfo(), name)
	}
}

func (p *parser) specialMeta() (map[string]int, error) {
	p.ws()
	var err error
	var meta map[string]int
	if p.peek(