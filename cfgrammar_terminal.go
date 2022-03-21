
package fmr

import "fmt"

func (p *parser) terminal() (flags, text string, err error) {
	if err = p.eat('"'); err != nil {
		return
	}
	p.ws()
	if p.peek() == '(' {
		p.eat('(')
		p.ws()
		if err = p.eat('?'); err != nil {
			return
		}
		p.ws()
		if flags, err = p.text(); err != nil {
			return
		}
		if err = p.eat(')'); err != nil {
			return
		}
	}
	if text, err = p.terminalText(); err != nil {
		return
	}
	err = p.eat('"')
	return
}

func (p *parser) terminalText() (string, error) {
	var ret []rune
	var prev rune
	for {
		switch r := p.next(); {
		case r == '"' && prev != '\\':
			p.backup()
			return string(ret), nil
		case r == eof:
			return "", fmt.Errorf("%s : unterminated string", p.posInfo())
		case prev == '\\':
			switch r {