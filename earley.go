
package fmr

import (
	"fmt"

	"github.com/liuzl/ling"
)

// GammaRule is the name of the special "gamma" rule added by the algorithm
// (this is unicode for 'LATIN SMALL LETTER GAMMA')
const GammaRule = "\u0263" // "\u0194"

// DOT indicates the current position inside a TableState
const DOT = "\u2022" // "\u00B7"

// TableState uses Earley's dot notation: given a production X → αβ,
// the notation X → α • β represents a condition in which α has already
// been parsed and β is expected.