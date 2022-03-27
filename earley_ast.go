
package fmr

import "fmt"

// Debug flag
var Debug = false

// Node is the AST of tree structure
type Node struct {
	Value    *TableState `json:"value"`
	Children []*Node     `json:"children,omitempty"`

	p *Parse
}

// GetFinalStates returns the final states of p
func (p *Parse) GetFinalStates() []*TableState {
	return p.finalStates
}

// Boundary returns the start, end position in NL for finalState
func (p *Parse) Boundary(finalState *TableState) *Pos {
	if finalState == nil {
		return nil
	}
	start := p.columns[finalState.Start+1].token.StartByte
	end := p.columns[finalState.End].token.EndByte
	if end < start { //TODO
		end = start
	}
	return &Pos{start, end}
}

// Tag returns the Nonterminal name of finalState
func (p *Parse) Tag(finalState *TableState) string {
	if finalState == nil {
		return ""
	}
	return finalState.Rb.Terms[0].Value
}

// GetTrees returns all possible parse results
func (p *Parse) GetTrees(finalState *TableState) []*Node {
	if Debug {
		fmt.Printf("chart:\n%+v\n", p)
		fmt.Println("finalState:\n", finalState)
	}
	if finalState != nil {
		return p.buildTrees(finalState)
	}
	return nil
}

func (p *Parse) buildTrees(state *TableState) []*Node {
	if state.Term.Type == Any {
		n := &TableState{state.Term, nil, state.Start, state.End, state.End}
		cld := []*Node{{n, nil, p}}
		return cld