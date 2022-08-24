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
	var