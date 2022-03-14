package fmr

import "fmt"

func (p *parser) special() (*Term, error) {
	if err := p.eat('('); err !=