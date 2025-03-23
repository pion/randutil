// SPDX-FileCopyrightText: 2023 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package randutil

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMathRandomGenerator(t *testing.T) {
	g := NewMathRandomGenerator()
	isLetter := regexp.MustCompile(`^[a-zA-Z]+$`).MatchString

	for i := 0; i < 10000; i++ {
		s := g.GenerateString(10, runesAlpha)
		assert.Equal(t, 10, len(s), "Generated string was not the correct length")
		assert.True(t, isLetter(s), "Generator returned unexpected character: %s", s)
	}
}

func TestIntn(t *testing.T) {
	g := NewMathRandomGenerator()

	localMin := 100
	localMax := 0
	for i := 0; i < 10000; i++ {
		r := g.Intn(100)
		assert.GreaterOrEqual(t, r, 0, "Generated value was not greater than 0")
		assert.Less(t, r, 100, "Generated value was not less than 100")

		if r < localMin {
			localMin = r
		}
		if r > localMax {
			localMax = r
		}
	}
	assert.Greater(t, 10, localMin, "Value around lower boundary was not generated")
	assert.Less(t, 90, localMax, "Value around upper boundary was not generated")
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
	assert.Greater(t, uint64(0x1000000000000000), localMin, "Value around lower boundary was not generated")
	assert.Less(t, uint64(0xF000000000000000), localMax, "Value around upper boundary was not generated")
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
	assert.Greater(t, uint32(0x10000000), localMin, "Value around lower boundary was not generated")
	assert.Less(t, uint32(0xF0000000), localMax, "Value around upper boundary was not generated")
}
