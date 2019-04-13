# LFSR [![Build Status](https://travis-ci.org/taylorza/go-lfsr.svg?branch=master)](https://travis-ci.org/taylorza/go-lfsr) [![Coverage Status](https://coveralls.io/repos/github/taylorza/go-lfsr/badge.svg?branch=master)](https://coveralls.io/github/taylorza/go-lfsr?branch=master) [![Go Report Card](https://goreportcard.com/badge/github.com/taylorza/go-lfsr)](https://goreportcard.com/report/github.com/taylorza/go-lfsr) [![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/taylorza/go-lfsr) 
Linear feedback shift register (LFSR) based psuedo random number generator

Random numbers generated by the LFSR implemented in this package will produce a sequence that has a period covering the full range of bits for the chosen LFSR.
A well desing n bit LFSR that will return the numbers in the range 1-2^n in a pseudo random order with no repeating numbers until the full range has been exhausted.
In this package the following linear feedback registers are provided:
Lfsr8 - 8 bit LFSR with a period of 255
Lfsr16 - 16 bit LFSR with a period of 65,535
Lfsr32 - 32 bit LFSR with a period of 4,294,967,296
Lfsr64 - 64 bit LFSR with a period of 18,446,744,073,709,551,616

## Installation

Use the 'go' command:

    $go get github.com/taylorza/go-lfsr

## Examples

```go
    // Create a 64 bit LFSR with a time based seed
    l := lfsr.NewLfsr64(0)

    // Get the next random number in the sequence from the LFSR
    v, _ := l.Next();
```

See [tests](https://github.com/taylorza/go-lfsr/blob/master/lfsr_test.go) for more examples

## Copyright 
(C)2013-2019 by Chris Taylor (taylorza)
See [Licence](https://github.com/taylorza/go-lfsr/blob/master/LICENSE)