package fmr

import (
	"fmt"
	"strings"

	"github.com/liuzl/ling"
	"github.com/liuzl/unidecode"
	"github.com/mitchellh/hashstructure"
)

func (g *Grammar) refine(prefix string) error {
	if g.Refined {
		return nil
	}
	var terminalRules []*Rule
	var terminals = make(map[string]string)
	var names = make(map[string]bool)
	var n int
	var name string
	for _, rule := range g.Rules {
		for _, body := range rule.Body {
			for _, term := range body.Terms {
				if term.Type != Terminal {
					continue
				}
				// if this is a terminal text inside a ruleBody
				if t, has := terminals[term.Value]; has {
					term.Value = t
				} else {
					d := ling.NewDocument(term.Value)
					if err := NLP().Annotate(d); err != nil {
						return err
					}
					tname := prefix + "_t"
					rb := &RuleBody{}
					for _, token := range d.Tokens {
		