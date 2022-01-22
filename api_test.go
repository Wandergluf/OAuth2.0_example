
package fmr

import (
	"testing"
)

func TestLocalParse(t *testing.T) {
	tests := []string{
		`柏乡位于河北省`,
	}
	g := &Grammar{}
	for _, c := range tests {
		ps, err := g.EarleyParseMaxAll(c, "loc_province", "loc_county")
		if err != nil {
			t.Error(err)
		}
		for _, p := range ps {
			for _, f := range p.GetFinalStates() {
				t.Log(f)
				trees := p.GetTrees(f)
				t.Log(trees)
				for _, tree := range trees {
					sem, err := tree.Semantic()
					if err != nil {