
package fmr

import (
	"fmt"

	"github.com/liuzl/ling"
)

// GammaRule is the name of the special "gamma" rule added by the algorithm
// (this is unicode for 'LATIN SMALL LETTER GAMMA')
const GammaRule = "\u0263" // "\u0194"

// DOT indicates the current position inside a TableState
const DOT = "\u2022" // "\u00B7"

// TableState uses Earley's dot notation: given a production X → αβ,
// the notation X → α • β represents a condition in which α has already
// been parsed and β is expected.
type TableState struct {
	Term  *Term     `json:"term"`
	Rb    *RuleBody `json:"rb,omitempty"`
	Start int       `json:"start"`
	End   int       `json:"end"`
	Dot   int       `json:"dot"`
}

// TableColumn is the TableState set
type TableColumn struct {
	token  *ling.Token
	index  int
	states []*TableState
}

// Parse stores a parse chart by grammars
type Parse struct {
	grammars    []*Grammar
	text        string
	starts      []string
	columns     []*TableColumn
	finalStates []*TableState
}

// Equal func for TableState
func (s *TableState) Equal(ts *TableState) bool {
	if s == nil && ts == nil {
		return true
	}
	if s == nil || ts == nil {
		if Debug {
			fmt.Println("only one is nil:", s, ts)
		}
		return false
	}
	if s.Start != ts.Start || s.End != ts.End || s.Dot != ts.Dot ||
		!s.Rb.Equal(ts.Rb) {
		return false
	}
	return s.Term.Equal(ts.Term)
}

func (s *TableState) metaEmpty() bool {
	if s.Term.Meta == nil {
		return true
	}
	if m, ok := s.Term.Meta.(map[string]int); ok && len(m) == 0 {
		return true
	}
	return false
}

func (s *TableState) isCompleted() bool {
	switch s.Term.Type {
	case Any, List:
		if !s.metaEmpty() {
			if meta, ok := s.Term.Meta.(map[string]int); ok {
				if s.Dot >= meta["min"] && s.Dot <= meta["max"] {
					return true
				}
			}
		} else if s.Dot > 0 {