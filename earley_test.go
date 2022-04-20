package fmr

import (
	"bytes"
	"testing"

	"zliu.org/goutil"
)

func TestEarleyParse(t *testing.T) {
	var grammar = `<expr> = "a" | "a" "+" <expr> {nf.math.sum($1, $3)};`
	//grammar = `<ex