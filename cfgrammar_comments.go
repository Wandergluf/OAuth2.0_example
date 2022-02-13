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
		if len(c) == 0 {
			return nil
		}
	}
}

func (p *parser) comment() (string, error) {
	if p.next() != '/' {
		p.backup()
		ret