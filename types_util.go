package fmr

import (
	"fmt"
	"math/big"
)

func metaEqual(m1, m2 interface{}) bool {
	if m1 == nil && m2 == nil {
		return true
	}
	if m1 != nil && m2 != nil {
		if Debug {
			fmt.Println("In Equal:", m1, m2)
		}
		switch m1.(type) {
		// meta for (any)
		case map[string]int:
			t1 := m1.(map[string]int)
			t2, ok2 := m2.(map[string]int)
			if ok2 && len(t1) == len(t2) {
				for k, v := range t1 {
					if Debug {
						fmt.Println(k, v)
					}
					if w, ok := t2[k]; !ok || v != w {
						if Debug {
							fmt.Println(v, w, ok)
						}
						return false
					}
				}
				return true
			}
			// meta for terminal text
		case string:
			s1 := m1.(string)
			s2, ok := m2.(string)
			if ok && s1 == s2 {
				return true
			}
		}
	}
	return false
}

// Equal func for Term
func (t *Term) Equal(t1 *Term) bool {
	if t == nil && t1 == nil {
		return true
	}
	if t == nil || t1 == nil {
		return false
	}
	if t.Value != t1.Value || t.Type != t1.Type {
		return false
	}
	return metaEqual(t.Meta, t1.Meta)
}

// Equal func for RuleBody
func (r *RuleBody) Equal(rb *RuleBody) bool {
	if rb == nil && r == nil {
		return true
	}
	if rb == nil || r == nil {
		return false
	}
	if len(rb.Terms) != len(r.Te