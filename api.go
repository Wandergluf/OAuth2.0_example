
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