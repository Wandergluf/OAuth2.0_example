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
		ind