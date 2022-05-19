package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/golang/glog"
	"github.com/liuzl/fmr"
	"github.com/robertkrimen/otto"
	"zliu.org/goutil"
)

var (
	grammar = flag.String("g", "builtin.grammar", "grammar file")
	js      = flag.String("js", "math.js", "javascript file")
	input   = flag.String("i", "", "file of original text to read")
	debug   = flag.Bool("debug", false, "debug mode")
	start   = flag.String("start", "number", "start rule")
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	flag.Parse()
	if *debug {
		fmr.Debug = true
	}
	g, err := fmr.GrammarFromFile(*grammar)
	if err != nil {
		glog.Fatal(err)
	}
	if *debug {
		b, err := goutil.JSONMarshalIndent