
package fmr

import (
	"fmt"
)

// FrameFMR parses NL text to FMR
func (g *Grammar) FrameFMR(text string) ([]string, error) {
	return g.FrameFMRWithContext("", text)
}

// FrameFMRWithContext parses NL text to FMR
func (g *Grammar) FrameFMRWithContext(context, text string) ([]string, error) {
	frames, err := g.MatchFramesWithContext(context, text)
	if err != nil {
		return nil, err
	}
	var ret []string
	for k, v := range frames {
		f := g.Frames[k.RuleName].Body[k.BodyID].F
		terms := g.Frames[k.RuleName].Body[k.BodyID].Terms
		var children []*Node
		for _, term := range terms {
			slots := v.Slots[term.Key()]
			if slots == nil || len(slots) == 0 || len(slots[0].Trees) == 0 {
				children = append(children, nil)
				continue
			}