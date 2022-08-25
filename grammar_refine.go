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
	for _, rule := range g.Rules 