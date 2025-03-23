// SPDX-FileCopyrightText: 2023 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package randutil

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCryptoRandomGenerator(t *testing.T) {
	isLetter := regexp.MustCompile(`^[a-zA-Z]+$`).MatchString

	for i := 0; i < 10000; i++ {
		s, err := GenerateCryptoRandomString(10, runesAlpha)
		assert.NoError(t, err)
		assert.Equal(t, 10, len(s), "Generated string was not the correct length")
		assert.True(t, isLetter(s), "Generator returned unexpected character: %s", s)
	}
}

func TestCryptoUint64(t *testing.T) {
	localMin := uint64(0xFFFFFFFFFFFFFFFF)
	localMax := uint64(0)
	for i := 0; i < 10000; i++ {
		r, err := CryptoUint64()
		assert.NoError(t, err)

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
