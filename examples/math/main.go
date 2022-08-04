
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"runtime/pprof"
	"strings"

	"github.com/golang/glog"
	"github.com/liuzl/fmr"
	"github.com/robertkrimen/otto"
)

var (
	grammar    = flag.String("g", "grammars/math.grammar", "grammar file")
	js         = flag.String("js", "math.js", "javascript file")
	input      = flag.String("i", "", "file of original text to read")
	start      = flag.String("start", "number", "start rule")
	eval       = flag.Bool("eval", false, "execute flag")
	debug      = flag.Bool("debug", false, "debug mode")