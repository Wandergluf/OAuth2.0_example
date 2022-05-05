
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/golang/glog"
	"github.com/liuzl/fmr"
	"github.com/robertkrimen/otto"
)

var (
	grammar = flag.String("g", "arithmetic.grammar", "grammar file")
	js      = flag.String("js", "math.js", "javascript file")
	input   = flag.String("i", "", "file of original text to read")
)

func main() {
	flag.Parse()
	//bnf.Debug = true
	g, err := fmr.GrammarFromFile(*grammar)
	if err != nil {
		glog.Fatal(err)
	}

	script, err := ioutil.ReadFile(*js)
	if err != nil {
		glog.Fatal(err)
	}
	vm := otto.New()
	if _, err = vm.Run(script); err != nil {
		glog.Fatal(err)
	}

	var in *os.File
	if *input == "" {
		in = os.Stdin
	} else {
		in, err = os.Open(*input)
		if err != nil {
			glog.Fatal(err)
		}
		defer in.Close()
	}
	br := bufio.NewReader(in)

	for {
		line, c := br.ReadString('\n')