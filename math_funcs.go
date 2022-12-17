
package fmr

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
)

func init() {
	builtinFuncs["nf.math.sum"] = sum
	builtinFuncs["nf.math.sub"] = sub
	builtinFuncs["nf.math.mul"] = mul
	builtinFuncs["nf.math.div"] = div
	builtinFuncs["nf.math.pow"] = pow
	builtinFuncs["nf.math.neg"] = neg
	builtinFuncs["nf.math.even"] = even
	builtinFuncs["nf.math.odd"] = odd
	builtinFuncs["nf.math.prime"] = prime
}

func sum(x, y string) string {
	return calc(x, y, "Add")
}

func sub(x, y string) string {