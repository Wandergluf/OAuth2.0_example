package fmr

import (
	"fmt"
	"strings"

	"github.com/liuzl/dict"
)

func updateIndex(index map[string]*Index, k string, cate string, v RbKey) error {
	if index == nil {
		return fmt.Errorf("nil grammar index")
	}
	if cate != "frame" && cate != "rule" {
		return fmt.Errorf("invalid cate %s", cate)
	}
	if index[k] == nil {
		index[k] = &Index{make(map[RbKey]struct{}), make(map[RbKey]struct{})}
	}
	switch cate {
	case "frame":
		index[k].Frames[v] = struct{}{}
	case "rule":
		index[k].Rules[v] = struct{}{}
	}
	return nil
}

func (g *Grammar) indexRules(rules map[string]*Rule, cate string) error {
	var err error
	for _, rule := range rules {
		for id, body := range rule.Body {
			for _, term := range body.Terms {
				v := RbKey{rule.Name, id}
				value := strings.TrimSpace(term.Value)
				if value == "" {
					continue
				}
				switch term.Type {
				case Terminal:
					if err = g.trie.SafeUpdate([]byte(value), 1); err != nil {
						return err
					}
					if err = updateIndex(g.index, value, cate, v); err != nil {
						return err
					}
				case No