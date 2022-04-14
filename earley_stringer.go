
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