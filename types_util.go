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
						fmt.Println