package fmr

import (
	"bytes"
	"testing"

	"zliu.org/goutil"
)

func TestEarleyParse(t *testing.T) {
	var grammar = `<expr> =