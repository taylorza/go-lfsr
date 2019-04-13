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
	"testing"
)

func TestLfsr8(t *testing.T) {
	p := 0
	l := NewLfsr8(0)
	restarted := false
	for !restarted {
		p++
		_, restarted = l.Next()

	}
	if p != 0xff {
		t.Error("Lfsr8 did not complete the full period and restarted at ", p)
	}
}
func BenchmarkLfsr8(b *testing.B) {
	l := NewLfsr8(0)
	for i := 0; i < b.N; i++ {
		_, _ = l.Next()
	}
}

func TestLfsr16(t *testing.T) {
	p := 0
	l := NewLfsr16(0)
	restarted := false
	for !restarted {
		p++
		_, restarted = l.Next()

	}
	if p != 0xffff {
		t.Error("Lfsr16 did not complete the full period and restarted at ", p)
	}
}

func BenchmarkLfsr16(b *testing.B) {
	l := NewLfsr16(0)
	for i := 0; i < b.N; i++ {
		_, _ = l.Next()
	}
}

func TestLfsr32(t *testing.T) {
	p := 0
	l := NewLfsr32(0)
	restarted := false
	for !restarted {
		p++
		_, restarted = l.Next()

	}
	if p != 0xffffffff {
		t.Error("Lfsr32 did not complete the full period and restarted at ", p)
	}
}

func BenchmarkLfsr32(b *testing.B) {
	l := NewLfsr32(0)
	for i := 0; i < b.N; i++ {
		_, _ = l.Next()
	}
}

func TestLfsr64(t *testing.T) {
	p := uint64(0)
	l := NewLfsr64(0)
	restarted := false
	for !restarted && p < 0x10000000 {
		p++
		_, restarted = l.Next()

	}
	if p != 0x10000000 {
		t.Error("Lfsr64 did not complete the full period and restarted at ", p)
	}
}

func BenchmarkLfsr64(b *testing.B) {
	l := NewLfsr64(0)
	for i := 0; i < b.N; i++ {
		_, _ = l.Next()
	}
}
