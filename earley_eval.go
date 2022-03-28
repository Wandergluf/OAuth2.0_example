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
	