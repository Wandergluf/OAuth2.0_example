
package fmr

import (
	"fmt"

	"github.com/liuzl/ling"
	"github.com/mitchellh/hashstructure"
)

func (g *Grammar) localGrammar(d *ling.Document) (*Grammar, error) {
	if d == nil {
		return nil, fmt.Errorf("document is empty")
	}
	g.regexpTag(d)
	if len(d.Spans) == 0 && len(d.Tokens) == 0 {
		return nil, nil
	}
	l := &Grammar{Name: "local", Rules: make(map[string]*Rule)}
	for _, token := range d.Tokens {
		k := ""
		switch token.Type {
		case ling.Word:
			k = "word"
		case ling.Punct:
			k = "punct"
		case ling.Symbol:
			k = "symbol"
		case ling.Letters:
			k = "letters"
		}