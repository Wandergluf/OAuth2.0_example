package fmr

import (
	"fmt"
)

func (p *parser) comments() error {
	defer p.ws()
	for {
		