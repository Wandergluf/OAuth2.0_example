package fmr

import (
	"sync"

	"github.com/liuzl/ling"
)

type cMap struct {
	tokens map[string]*ling.