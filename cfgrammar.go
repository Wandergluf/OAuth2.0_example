
package fmr

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/mitchellh/hashstructure"
)

type parser struct {
	input   string
	pos     int
	width   int
	current *position
	info    map[int]*position
	fname   string
	dir     string
}

type position struct {
	row, col int
	r        string
}

func (p *position) String() string {
	return fmt.Sprintf("|row:%d, col:%d, c:%s|", p.row, p.col, strconv.Quote(p.r))
}

const eof = -1

// GrammarFromFile constructs the Context-Free Grammar from file
func GrammarFromFile(file string) (*Grammar, error) {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	dir, err := filepath.Abs(filepath.Dir(file))
	if err != nil {
		return nil, err
	}
	return grammarFromString(string(b), file, dir, map[string]int{file: 1})
}

func grammarFromFile(ifile string, files map[string]int) (*Grammar, error) {
	if files[ifile] >= 2 {
		return nil, nil
	}
	b, err := ioutil.ReadFile(ifile)
	if err != nil {
		return nil, err
	}
	dir, err := filepath.Abs(filepath.Dir(ifile))
	if err != nil {
		return nil, err
	}
	return grammarFromString(string(b), ifile, dir, files)
}

// GrammarFromString constructs the Contex-Free Grammar from string d with name
func GrammarFromString(d, name string) (*Grammar, error) {
	return grammarFromString(d, name, ".", make(map[string]int))
}

func grammarFromString(d, name, dir string, files map[string]int) (*Grammar, error) {
	if files[name] >= 2 {
		return nil, nil
	}
	p := &parser{fname: name, dir: dir, input: d, info: make(map[int]*position)}
	if Debug {
		fmt.Println("loading ", name, files)
	}
	g, err := p.grammar(files)
	if err != nil {
		return nil, err
	}
	files[name]++
	if Debug {
		fmt.Println("loaded ", name, files)
	}
	return g, nil
}

func (p *parser) posInfo() string {
	return fmt.Sprintf("%s%s", p.fname, p.current)
}

func (p *parser) next() rune {
	if p.pos >= len(p.input) {
		p.width = 0
		return eof
	}
	r, w := utf8.DecodeRuneInString(p.input[p.pos:])
	if r == utf8.RuneError {
		return eof
	}
	p.width = w
	p.pos += w
	if p.info[p.pos] == nil {
		if p.current == nil {
			p.current = &position{1, w, string(r)}
		} else {
			if r == '\n' {
				p.current = &position{p.current.row + 1, w, string(r)}
			} else {
				p.current = &position{p.current.row, p.current.col + w, string(r)}
			}
		}
		p.info[p.pos] = p.current
	} else {
		p.current = p.info[p.pos]
	}
	return r
}

func (p *parser) eat(expected rune) error {
	if r := p.next(); r != expected {
		return fmt.Errorf("%s :expected %s, got %s", p.posInfo(),
			strconv.Quote(string(expected)), strconv.Quote(string(r)))
	}
	return nil
}

func (p *parser) backup() {
	p.pos -= p.width
	p.current = p.info[p.pos]
}

func (p *parser) peek() rune {
	r := p.next()
	p.backup()
	return r
}
