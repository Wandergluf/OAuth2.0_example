
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
		if k != "" {
			rb := &RuleBody{
				[]*Term{{Value: token.String(), Type: Terminal}}, nil}
			hash, err := hashstructure.Hash(rb, nil)
			if err != nil {
				return nil, err
			}
			if _, has := l.Rules[k]; has {
				l.Rules[k].Body[hash] = rb
			} else {
				l.Rules[k] = &Rule{k, map[uint64]*RuleBody{hash: rb}}
			}
		}
	}
	for _, span := range d.Spans {
		if span.Annotations["value"] == nil {
			continue
		}
		m, ok := span.Annotations["value"].(map[string]interface{})
		if !ok {
			continue
		}
		//terms := []*Term{{Value: span.String(), Type: Terminal}}
		var terms []*Term
		for i := span.Start; i < span.End; i++ {
			terms = append(terms,
				&Term{Value: span.Doc.Tokens[i].Text, Type: Terminal})
		}
		for k, values := range m {
			rb := &RuleBody{terms, nil}
			switch values.(type) {
			case []interface{}:
				args := []*Arg{}
				for _, v := range values.([]interface{}) {
					args = append(args, &Arg{"string", v})
				}
				rb.F = &FMR{"fmr.entity",
					[]*Arg{{"string", k}, {"func", &FMR{"fmr.list", args}}},
				}
				rb.F = &FMR{"fmr.list",
					[]*Arg{{"string", span.String()}, {"func", rb.F}}}
			}
			hash, err := hashstructure.Hash(rb, nil)
			if err != nil {
				return nil, err
			}
			if _, has := l.Rules[k]; has {
				l.Rules[k].Body[hash] = rb
			} else {
				l.Rules[k] = &Rule{k, map[uint64]*RuleBody{hash: rb}}
			}
		}
	}
	if len(l.Rules) == 0 {
		return nil, nil
	}
	/*
		if err := l.refine("l"); err != nil {
			return nil, err
		}
	*/
	return l, nil