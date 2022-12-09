
package fmr

import (
	"bytes"
	"testing"

	"zliu.org/goutil"
)

func TestList(t *testing.T) {
	//Debug = true
	cases := []string{
		`直辖市：北京上海天津`,
		`直辖市：北京、上海和天津`,
		`直辖市：北京、上海和天津、津城`,
		`直辖市：帝都、魔都、寨都、旧都`,
		`直辖市：北京`,
	}
	g, err := GrammarFromFile("sf.grammar")