package fmr

import (
	//"fmt"
	"testing"

	"zliu.org/goutil"
)

var tests = []string{
	`<list>  =  "<" <items> ">"               ;
	<items> =  <items> " " <item> {     nf.math.sum($1,$3