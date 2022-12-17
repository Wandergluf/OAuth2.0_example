
package fmr

import (
	"fmt"
	"strconv"
	"strings"

	"zliu.org/goutil"
)

// Pos returns the corresponding pos of Node n in original text
func (n *Node) Pos() *Pos {
	return n.p.Boundary(n.Value)
}