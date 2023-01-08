
package fmr

import (
	"encoding/gob"
	"fmt"

	"github.com/liuzl/dict"
	"github.com/mitchellh/hashstructure"
)

func init() {
	gob.Register(RbKey{})
}

// A Grammar stores a Context-Free Grammar
type Grammar struct {
	Name    string            `json:"name"`
	Rules   map[string]*Rule  `json:"rules"`
	Frames  map[string]*Rule  `json:"frames"`
	Regexps map[string]string `json:"regexps"`
	Refined bool              `json:"refined"`

	trie      *dict.Cedar
	index     map[string]*Index
	ruleIndex map[string]*Index

	includes []*Grammar
}

// An Index contains two sets for frames' names and rules' names
type Index struct {
	Frames map[RbKey]struct{}
	Rules  map[RbKey]struct{}
}

// A RbKey identifies a specific RuleBody by name and id
type RbKey struct {
	RuleName string `json:"rule_name"`
	BodyID   uint64 `json:"body_id"`
}

// A Pos specifies the start and end positions
type Pos struct {
	StartByte int `json:"start_byte"`
	EndByte   int `json:"end_byte"`
}

// A Slot contains the Pos and its corresponding parse trees
type Slot struct {
	Pos
	Trees []*Node
}

// A Frame is a frame consists of Slots
type Frame struct {
	Slots    map[uint64][]*Slot
	Complete bool
}

func (f *Frame) String() string {
	return fmt.Sprintf("Complete:%+v, %+v", f.Complete, f.Slots)
}
