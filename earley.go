
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
			return true
		}
		return false
	default:
		return s.Dot >= len(s.Rb.Terms)
	}
}

func (s *TableState) getNextTerm() *Term {
	switch s.Term.Type {
	case Any:
		if !s.metaEmpty() {
			if meta, ok := s.Term.Meta.(map[string]int); ok && s.Dot >= meta["max"] {
				return nil
			}
		}
		return s.Term
	case List:
		if !s.metaEmpty() {
			if meta, ok := s.Term.Meta.(map[string]int); ok && s.Dot >= meta["max"] {
				return nil
			}
		}
		return &Term{Value: s.Term.Value, Type: Nonterminal, Meta: s.Term.Meta}
	default:
		if s.isCompleted() {
			return nil
		}
		return s.Rb.Terms[s.Dot]
	}
}

func (col *TableColumn) insert(state *TableState) *TableState {
	return col.insertToEnd(state, false)
}

func (col *TableColumn) insertToEnd(state *TableState, end bool) *TableState {
	state.End = col.index
	if state.Term.Type == Any {
		state.Dot = state.End - state.Start
	}
	for i, s := range col.states {
		if s.Equal(state) {
			if end {
				col.states = append(col.states[:i], col.states[i+1:]...)
				col.states = append(col.states, s)
			}
			return s
		}
	}
	col.states = append(col.states, state)
	return col.states[len(col.states)-1]
}

/*
 * the Earley algorithm's core: add gamma rule, fill up table, and check if the
 * gamma rule span from the first column to the last one. return the final gamma
 * state, or null, if the parse failed.
 */
func (p *Parse) parse(maxFlag bool) []*TableState {
	if len(p.starts) == 0 {