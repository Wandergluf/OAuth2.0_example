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
		return "", nil
	}
	switch r := p.peek(); {
	case r == '/':
		return p.lineComment()
	case r == '*':
		return p.multiLineComment()
	default:
		return "", fmt.Errorf("%s : invalid char %s", p.posInfo(), string(r))
	}
}

func (p *