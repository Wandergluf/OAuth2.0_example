package fmr

import (
	"strings"

	"github.com/liuzl/ling"
)

func terminalMatch(term *Term, token *ling.Token) bool {
	if term == nil || token == nil || term.Type != Terminal {
		return false
	}
	t := gTokens.get(term.Value)
	if term.Meta == nil || t == nil {
		if term.Valu