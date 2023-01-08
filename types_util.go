package fmr

import (
	"fmt"
	"math/big"
)

func metaEqual(m1, m2 interface{}) bool {
	if m1 == nil && m2 == nil {
