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
			return "", fmt.Errorf("the length of Args of nf.I should be one")
		}
		s, err := n.semEval(f.Args[0], children)
		if err != nil {
			return "", err
		}
		return s, nil
	}

	var args []interface{}
	for _, arg := range f.Args {
		s, err := n.semEval(arg, children)
		if err != nil {
			return "", err
		}
		args = append(args, s)
	}
	if Debug {
		fmt.Pri