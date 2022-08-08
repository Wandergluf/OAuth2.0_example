
package fmr

import (
	"testing"
)

func TestMatchFrames(t *testing.T) {
	cases := []string{
		`从北京飞上海`,
		`飞上海，从北京，后天`,