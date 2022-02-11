package fmr

import (
	"fmt"
)

func (p *parser) comments() error {
	defer p.ws()
	for {
		p.ws()
		c, err := p.comment()
		if err != nil {
			return err
		}