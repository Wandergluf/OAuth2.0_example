package fmr

import (
	"sync"

	"github.com/liuzl/ling"
)

type cMap struct {
	tokens map[string]*ling.Token
	sync.RWMutex
}

func (m *cMap) get(k string) *ling.Token {
	m.RLock()
	defer m.RUnlock()
	ret