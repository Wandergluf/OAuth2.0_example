package fmr

import (
	"github.com/liuzl/ling"
	"zliu.org/goutil"
)

func (g *Grammar) regexpTag(d *ling.Document) {
	if d == nil || len(d.Tokens) == 0 || len(g.Regexps) == 0 {
		return
	}

	for typ, s := range