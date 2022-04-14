
package fmr

import (
	"fmt"
	"io"
	"math/big"
	"strconv"
	"strings"
)

func (ts *TableState) String() string {
	s := ""
	switch ts.Term.Type {
	case Nonterminal:
		if ts.Rb != nil {
			for i, term := range ts.Rb.Terms {
				if i == ts.Dot {
					s += DOT + " "
				}
				switch term.Type {
				case Nonterminal:
					s += "<" + term.Value + "> "
				case Terminal:
					s += strconv.Quote(term.Value) + " "
				case Any:
					s += "(any) "
				case List:
					s += "(list<" + term.Value + ">) "
				}
			}
			if ts.Dot == len(ts.Rb.Terms) {
				s += DOT
			}
			return fmt.Sprintf("<%s> -> %s [%d-%d] {%s}",
				ts.Term.Value, s, ts.Start, ts.End, ts.Rb.F)
		}
	case Any:
		for i := ts.Start; i < ts.End; i++ {
			s += "# "
		}
		s += DOT + " * "
		return fmt.Sprintf("(any) -> %s [%d-%d]", s, ts.Start, ts.End)
	case List:
		f := "fmr.list("
		for i := 0; i < ts.Dot; i++ {
			s += "<" + ts.Term.Value + "> "
			f += fmt.Sprintf("$%d", i+1)
			if i != ts.Dot-1 {
				f += ","
			}
		}
		f += ")"
		s += DOT + " * "
		return fmt.Sprintf("(list<%s>) -> %s [%d-%d] {%s}", ts.Term.Value, s, ts.Start, ts.End, f)
	}
	return fmt.Sprintf("%s [%d-%d]", strconv.Quote(ts.Term.Value), ts.Start, ts.End)
}

func (tc *TableColumn) String() (out string) {
	if tc.index == 0 {
		out = "[0] ''\n"
	} else {
		out = fmt.Sprintf("[%d] '%s' position:[%d-%d]\n",
			tc.index, tc.token, tc.token.StartByte, tc.token.EndByte)
	}
	out += "=======================================\n"
	for _, s := range tc.states {
		out += s.String() + "\n"
	}
	return out
}

func (p *Parse) String() string {
	out := ""
	for _, c := range p.columns {
		out += c.String() + "\n"
	}
	return out
}

// Print this tree to out
func (n *Node) Print(out io.Writer) {
	n.printLevel(out, 0)
}

func (n *Node) printLevel(out io.Writer, level int) {
	indentation := ""
	for i := 0; i < level; i++ {
		indentation += "  "
	}
	fmt.Fprintf(out, "%s%v\n", indentation, n.Value)
	for _, child := range n.Children {
		child.printLevel(out, level+1)
	}
}

func (n *Node) String() string {
	if len(n.Children) > 0 {
		return fmt.Sprintf("%+v %+v", n.Value, n.Children)
	}
	return fmt.Sprintf("%+v", n.Value)
}

func (f *FMR) String() string {
	if f == nil {
		return "nf.I($0)"
	}
	var args []string
	invalid := "invalid_fmr"
	for _, arg := range f.Args {
		switch arg.Type {
		case "string":
			if s, ok := arg.Value.(string); ok {
				args = append(args, strconv.Quote(s))
			} else {
				return invalid
			}
		case "int":