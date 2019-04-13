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
	for !restarted && p < 0x100000000 {
		p++
		_, restarted = l.Next()

	}
	if p != 0x100000000 {
		t.Error("Lfsr64 did not complete the full period and restarted at ", p)
	}
}

func BenchmarkLfsr64(b *testing.B) {
	l := NewLfsr64(0)
	for i := 0; i < b.N; i++ {
		_, _ = l.Next()
	}
}
