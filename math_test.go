// SPDX-FileCopyrightText: 2023 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package randutil

import (
	"regexp"
	"testing"
)

func TestMathRandomGenerator(t *testing.T) {
	g := NewMathRandomGenerator()
	isLetter := regexp.MustCompile(`^[a-zA-Z]+$`).MatchString

	for i := 0; i < 10000; i++ {
		s := g.GenerateString(10, runesAlpha)
		if len(s) != 10 {
			t.Error("Generator returned invalid length")
		}
		if !isLetter(s) {
			t.Errorf("Generator returned unexpected character: %s", s)
		}
	}
}

func TestIntn(t *testing.T) {
	g := NewMathRandomGenerator()

	localMin := 100
	localMax := 0
	for i := 0; i < 10000; i++ {
		r := g.Intn(100)
		if r < 0 || r >= 100 {
			t.Fatalf("Out of range of Intn(100): %d", r)
		}
		if r < localMin {
			localMin = r
		}
		if r > localMax {
			localMax = r
		}
	}
	if localMin > 10 {
		t.Error("Value around lower boundary was not generated")
	}
	if localMax < 90 {
		t.Error("Value around upper boundary was not generated")
	}
}

func TestUint64(t *testing.T) {
	g := NewMathRandomGenerator()

	localMin := uint64(0xFFFFFFFFFFFFFFFF)
	localMax := uint64(0)
	for i := 0; i < 10000; i++ {
		r := g.Uint64()
		if r < localMin {
			localMin = r
		}
		if r > localMax {
			localMax = r
		}
	}
	if localMin > 0x1000000000000000 {
		t.Error("Value around lower boundary was not generated")
	}
	if localMax < 0xF000000000000000 {
		t.Error("Value around upper boundary was not generated")
	}
}

func TestUint32(t *testing.T) {
	g := NewMathRandomGenerator()

	localMin := uint32(0xFFFFFFFF)
	localMax := uint32(0)
	for i := 0; i < 10000; i++ {
		r := g.Uint32()
		if r < localMin {
			localMin = r
		}
		if r > localMax {
			localMax = r
		}
	}
	if localMin > 0x10000000 {
		t.Error("Value around lower boundary was not generated")
	}
	if localMax < 0xF0000000 {
		t.Error("Value around upper boundary was not generated")
	}
}
