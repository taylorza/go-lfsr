// MIT License
//
// Copyright (c) 2019 Chris Taylor (taylorza)
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package lfsr

import (
	"time"
)

// Lfsr8 represents an 8 bit linear feedback shift register
type Lfsr8 struct {
	state uint8
	seed  uint8
}

// NewLfsr8 returns a linear feedback shift register initialized with the specified seed. If the seed is zero the seed is initialized using the current time.
func NewLfsr8(seed uint8) *Lfsr8 {
	for seed == 0 {
		seed = uint8(time.Now().Nanosecond() & 0xff)
	}
	return &Lfsr8{seed, seed}
}

// Next returns the next pseudo random number from the linear feedback shift register and the restarted flag
// which indicates that the sequence has completed and is restarting.
func (l *Lfsr8) Next() (value uint8, restarted bool) {
	s := l.state
	b := (s >> 0) ^ (s >> 2) ^ (s >> 3) ^ (s >> 4)
	l.state = (s >> 1) | (b << 7)
	return l.state, l.state == l.seed
}

// Lfsr16 represents an 16 bit linear feedback shift register
type Lfsr16 struct {
	state uint16
	seed  uint16
}

// NewLfsr16 returns a linear feedback shift register initialized with the specified seed. If the seed is zero the seed is initialized using the current time.
func NewLfsr16(seed uint16) *Lfsr16 {
	for seed == 0 {
		seed = uint16(time.Now().Nanosecond() & 0xffff)
	}
	return &Lfsr16{seed, seed}
}

// Next returns the next pseudo random number from the linear feedback shift register and the restarted flag
// which indicates that the sequence has completed and is restarting.
func (l *Lfsr16) Next() (value uint16, restarted bool) {
	s := l.state
	b := (s >> 0) ^ (s >> 2) ^ (s >> 3) ^ (s >> 5)
	l.state = (s >> 1) | (b << 15)
	return l.state, l.state == l.seed
}

// Lfsr32 represents an 32 bit linear feedback shift register
type Lfsr32 struct {
	state uint32
	seed  uint32
}

// NewLfsr32 returns a linear feedback shift register initialized with the specified seed. If the seed is zero the seed is initialized using the current time.
func NewLfsr32(seed uint32) *Lfsr32 {
	for seed == 0 {
		seed = uint32(time.Now().Nanosecond() & 0xffffffff)
	}
	return &Lfsr32{seed, seed}
}

// Next returns the next pseudo random number from the linear feedback shift register and the restarted flag
// which indicates that the sequence has completed and is restarting.
func (l *Lfsr32) Next() (value uint32, restarted bool) {
	s := l.state
	b := (s >> 0) ^ (s >> 2) ^ (s >> 6) ^ (s >> 7)
	l.state = (s >> 1) | (b << 31)
	return l.state, l.state == l.seed
}

// Lfsr64 represents a 64 bit linear feedback shift register
type Lfsr64 struct {
	state uint64
	seed  uint64
}

// NewLfsr64 returns a linear feedback shift register initialized with the specified seed. If the seed is zero the seed is initialized using the current time.
func NewLfsr64(seed uint64) *Lfsr64 {
	for seed == 0 {
		seed = uint64(time.Now().Nanosecond() & 0xffffffff)
	}
	return &Lfsr64{seed, seed}
}

// Next returns the next pseudo random number from the linear feedback shift register and the restarted flag
// which indicates that the sequence has completed and is restarting.
func (l *Lfsr64) Next() (value uint64, restarted bool) {
	s := l.state
	b := (s >> 0) ^ (s >> 1) ^ (s >> 3) ^ (s >> 4)
	l.state = (s >> 1) | (b << 63)
	return l.state, l.state == l.seed
}
