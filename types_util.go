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
		// meta f