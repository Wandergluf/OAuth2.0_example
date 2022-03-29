package fmr

import (
	"fmt"
	"math/big"
)

// Eval returns the denotation of Node n
func (n *Node) Eval() (interface{}, error) {
	if n.Value.Rb == nil || n.Value.Rb.F == nil {
		if n.p == nil {
			return "", nil
		}
		return n.OriginalText(), nil
	}
	return n.fmrEval(n.Value.Rb.F, n.Children)
}

func (n *Node) fmrEval(f *FMR, children []*Node) (interface{}, error) {
	if f == nil {
		return "", nil
	}
	if f.Fn == "nf.I" {
		if len(f.Args) != 1 {
			return "", fmt.Errorf("the length of