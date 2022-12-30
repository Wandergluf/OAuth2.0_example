package fmr

import (
	"github.com/liuzl/ling"
	"zliu.org/goutil"
)

func (g *Grammar) regexpTag(d *ling.Document) {
	if d == nil || len(d.Tokens) == 0 || len(g.Regexps) == 0 {
		return
	}

	for typ, s := range g.Regexps {
		re, err := goutil.Regexp(s)
		if err != nil {
			continue
		}
		matches := re.FindAllStringIndex(d.Text, -1)
		for _, match := range matches {
			start := -1
			end := -1
			for _, token := range d.Tokens {
				if token.StartByte == match[0] {
					start = token.I
				}
				if token.EndByte == match[1] {
					end =