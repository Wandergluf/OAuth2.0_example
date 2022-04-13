
package fmr

import (
	"encoding/json"
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

// Semantic returns the stringified FMR of Node n
func (n *Node) Semantic() (string, error) {
	nl := strconv.Quote(n.NL())
	if n.Value.Rb == nil || n.Value.Rb.F == nil {
		if n.p == nil {
			return "", nil
		}
		// by default, returns nf.I($0)
		return nl, nil
	}
	return n.fmrStr(n.Value.Rb.F, n.Children, nl)
}

func (n *Node) fmrStr(f *FMR, children []*Node, nl string) (string, error) {
	if f == nil {
		return "", nil
	}
	if f.Fn == "nf.I" {
		if len(f.Args) != 1 {
			return "", fmt.Errorf("the length of Args of nf.I should be one")
		}
		s, err := n.semStr(f.Args[0], children, nl)
		if err != nil {
			return "", err
		}
		return s, nil
	}

	var args []string
	for _, arg := range f.Args {
		s, err := n.semStr(arg, children, nl)
		if err != nil {
			return "", err
		}
		args = append(args, s)
	}
	return fmt.Sprintf("%s(%s)", f.Fn, strings.Join(args, ", ")), nil
}

func (n *Node) semStr(arg *Arg, nodes []*Node, nl string) (string, error) {
	if arg == nil {
		return "", fmt.Errorf("arg is nil")
	}
	switch arg.Type {
	case "string":
		if s, ok := arg.Value.(string); ok {
			return strconv.Quote(s), nil
		}
		return "", fmt.Errorf("arg.Value: %+v is not string", arg.Value)
	case "int":
		if i, ok := arg.Value.(*big.Int); ok {
			return i.String(), nil
		}
		return "", fmt.Errorf("arg.Value: %+v is not int", arg.Value)
	case "float":
		if f, ok := arg.Value.(*big.Float); ok {
			return f.String(), nil
		}
		return "", fmt.Errorf("arg.Value: %+v is not float", arg.Value)
	case "func":