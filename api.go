
package fmr

import (
	"flag"
	"fmt"
	"net/url"
	"strings"
	"sync"

	"github.com/liuzl/ling"
)

var (
	apiTagger = flag.String("api_tagger", "", "http address of api tagger")
	ctxTagger = flag.String("ctx_tagger", "", "http address of context tagger")
)

var nlp *ling.Pipeline
var once sync.Once

// NLP returns handler for the ling nlp toolkit
func NLP() *ling.Pipeline {
	once.Do(func() {
		var err error
		var tagger *ling.DictTagger
		if nlp, err = ling.DefaultNLP(); err != nil {
			panic(err)
		}
		if tagger, err = ling.NewDictTagger(); err != nil {
			panic(err)
		}
		if err = nlp.AddTagger(tagger); err != nil {
			panic(err)
		}
		if *apiTagger == "" {
			return
		}
		var tagger1 *ling.APITagger
		if tagger1, err = ling.NewAPITagger(*apiTagger); err != nil {
			panic(err)
		}
		if err = nlp.AddTagger(tagger1); err != nil {
			panic(err)
		}
	})
	return nlp
}

// EarleyParse parses text for rule <start> at beginning
func (g *Grammar) EarleyParse(text string, starts ...string) (*Parse, error) {
	return g.EarleyParseWithContext("", text, starts...)
}

// EarleyParseWithContext with context information
func (g *Grammar) EarleyParseWithContext(
	context, text string, starts ...string) (*Parse, error) {
	tokens, l, err := g.process(context, text)
	if err != nil {
		return nil, err
	}
	return g.earleyParse(true, text, tokens, l, starts...)
}

// EarleyParseAny parses text for rule <start> at any position
func (g *Grammar) EarleyParseAny(
	text string, starts ...string) (*Parse, error) {

	return g.EarleyParseAnyWithContext("", text, starts...)
}

//EarleyParseAnyWithContext with context information
func (g *Grammar) EarleyParseAnyWithContext(
	context, text string, starts ...string) (*Parse, error) {

	tokens, l, err := g.process(context, text)
	if err != nil {
		return nil, err
	}
	var p *Parse
	for i := 0; i < len(tokens); i++ {
		if p, err = g.earleyParse(
			true, text, tokens[i:], l, starts...); err != nil {
			return nil, err
		}
		if p.finalStates != nil {
			return p, nil
		}
	}
	return p, nil
}

// EarleyParseMaxAll extracts all submatches in text for rule <start>
func (g *Grammar) EarleyParseMaxAll(
	text string, starts ...string) ([]*Parse, error) {
	return g.EarleyParseMaxAllWithContext("", text, starts...)
}

// EarleyParseMaxAllWithContext with context information
func (g *Grammar) EarleyParseMaxAllWithContext(
	context, text string, starts ...string) ([]*Parse, error) {
	tokens, l, err := g.process(context, text)
	if err != nil {
		return nil, err
	}
	var ret []*Parse
	for i := 0; i < len(tokens); {
		p, err := g.earleyParse(true, text, tokens[i:], l, starts...)
		if err != nil {
			return nil, err
		}
		if p.finalStates != nil {
			ret = append(ret, p)
			max := 0
			for _, finalState := range p.finalStates {
				if finalState.End > max {
					max = finalState.End
				}
			}
			i += max
		} else {
			i++
		}
	}
	return ret, nil
}

// EarleyParseAll extracts all submatches in text for rule <start>
func (g *Grammar) EarleyParseAll(
	text string, starts ...string) ([]*Parse, error) {
	return g.EarleyParseAllWithContext("", text, starts...)
}
