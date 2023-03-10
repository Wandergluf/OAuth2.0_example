
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
			children = append(children, slots[0].Trees[0])
		}
		n := &Node{}
		str, err := n.fmrStr(f, children, "")
		if err != nil {
			return nil, err
		}
		ret = append(ret, str)
	}
	return ret, nil
}

// MatchFrames returns the matched frames for NL text
func (g *Grammar) MatchFrames(text string) (map[RbKey]*Frame, error) {
	return g.MatchFramesWithContext("", text)
}

// MatchFramesWithContext returns the matched frames for NL text
func (g *Grammar) MatchFramesWithContext(
	context, text string) (map[RbKey]*Frame, error) {
	frames, starts, err := g.getCandidates(text)
	if err != nil {
		return nil, err
	}
	ps, err := g.EarleyParseAllWithContext(context, text, starts...)
	if err != nil {
		return nil, err
	}
	for _, p := range ps {
		for _, finalState := range p.finalStates {
			tag := p.Tag(finalState)
			pos := p.Boundary(finalState)
			trees := p.GetTrees(finalState)

			if tag == "" || pos == nil {
				return nil, fmt.Errorf("invalid parse")
			}

			slot := &Slot{*pos, trees}

			ret := g.ruleIndex[tag]
			if ret == nil {
				continue
			}
			for rbKey := range ret.Frames {
				if frames[rbKey] == nil {
					frames[rbKey] = &Frame{make(map[uint64][]*Slot), false}
				}
				t := Term{Value: tag, Type: Nonterminal}
				frames[rbKey].Slots[t.Key()] = append(frames[rbKey].Slots[t.Key()], slot)
				if len(frames[rbKey].Slots) >=
					len(g.Frames[rbKey.RuleName].Body[rbKey.BodyID].Terms) {
					frames[rbKey].Complete = true
				}
			}
		}
	}
	return frames, nil
}

func (g *Grammar) getCandidates(text string) (
	map[RbKey]*Frame, []string, error) {

	matches, err := g.trie.MultiMatch(text)
	if err != nil {
		return nil, nil, err
	}
	frames := map[RbKey]*Frame{}
	rules := map[string]bool{}
	for word, hits := range matches {
		v := g.index[word]
		if v == nil {
			return nil, nil, fmt.Errorf("%s in trie but not in index", word)
		}
		for rbKey := range v.Frames {
			if frames[rbKey] == nil {
				frames[rbKey] = &Frame{make(map[uint64][]*Slot), false}
			}
			t := Term{Value: word, Type: Terminal}
			for _, hit := range hits {
				frames[rbKey].Slots[t.Key()] = append(frames[rbKey].Slots[t.Key()],
					&Slot{Pos{hit.StartByte, hit.EndByte}, nil})
			}
			if len(frames[rbKey].Slots) >=
				len(g.Frames[rbKey.RuleName].Body[rbKey.BodyID].Terms) {
				frames[rbKey].Complete = true
			}
		}
		for rbKey := range v.Rules {
			rules[rbKey.RuleName] = true
		}
	}
	var ruleList []string
	for k := range rules {
		ruleList = append(ruleList, k)
	}
	for {
		if len(ruleList) == 0 {
			break
		}
		r := ruleList[0]
		ruleList = ruleList[1:]

		ret := g.ruleIndex[r]
		if ret == nil {
			continue
		}
		for rbKey := range ret.Rules {
			if !rules[rbKey.RuleName] {
				ruleList = append(ruleList, rbKey.RuleName)
				rules[rbKey.RuleName] = true
			}
		}
	}
	var starts []string
	for k := range rules {
		starts = append(starts, k)
	}
	return frames, starts, nil
}