package fmr

import "fmt"

func (p *parser) special() (*Term, error) {
	if err := p.eat('('); err != nil {
		return nil, err
	}
	p.ws()
	name, err := p.text()
	if 