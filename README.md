# FMR: Functional Meaning Representation & Semantic Parsing Framework
[![GoDoc](https://godoc.org/github.com/liuzl/fmr?status.svg)](https://godoc.org/github.com/liuzl/fmr)[![Go Report Card](https://goreportcard.com/badge/github.com/liuzl/fmr)](https://goreportcard.com/report/github.com/liuzl/fmr)

## Projects that uses FMR

### mathsolver
* codes: https://github.com/liuzl/mathsolver
* demo: https://mathsolver.zliu.org/

## What is semantic parsing?
Semantic parsing is the process of mapping a natural language sentence into an intermediate logical form which is a formal representation of its meaning.

The formal representation should be a detailed representation of the complete meaning of the natural language sentence in a fully formal language that:

* Has a rich ontology of types, properties, and relations.
* Supports automated reasoning or execution.

## Representation languages
Early semantic parsers used highly domain-specific meaning representation languages, with later systems using more extensible languages like Prolog, lambda calculus, lambda dependancy-based compositional semantics (λ-DCS), SQL, Python, Java, and the Alexa Meaning Representation Language. Some work has used more exotic meaning representations, like query graphs or vector representations.

### FMR, a formal meaning representation language
* FMR stands for  functional meaning representation
* Context-Free Grammar for bridging NL and FMR
* *[VIM Syntax highlighting for FMR grammar file](https://github.com/liuzl/vim-fmr)*

## Tasks
* Grammar checkers
* Dialogue management
* Question answering
* Information extra