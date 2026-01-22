// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package randutil

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

const runesAlpha = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func TestRandomGeneratorCollision(t *testing.T) {
	g := NewMathRandomGenerator()

	testCases := map[string]struct {
		gen func(t *testing.T) string
	}{
		"MathRandom": {
			gen: func(*testing.T) string {
				return g.GenerateString(10, runesAlpha)
			},
		},
		"CryptoRandom": {
			gen: func(t *testing.T) string {
				t.Helper()

				s, err := GenerateCryptoRandomString(10, runesAlpha)
				assert.NoError(t, err)

				return s
			},
		},
	}

	const maxIterations = 100
	const iteration = 100

	for name, testCase := range testCases {
		testCase := testCase
		t.Run(name, func(t *testing.T) {
			for iter := 0; iter < iteration; iter++ {
				var wg sync.WaitGroup
				var mu sync.Mutex

				rands := make([]string, 0, maxIterations)

				for i := 0; i < maxIterations; i++ {
					wg.Add(1)
					go func() {
						r := testCase.gen(t)
						mu.Lock()
						rands = append(rands, r)
						mu.Unlock()
						wg.Done()
					}()
				}
				wg.Wait()

				assert.Equal(t, maxIterations, len(rands), "Failed to generate all randoms")

				for i := 0; i < maxIterations; i++ {
					for j := i + 1; j < maxIterations; j++ {
						assert.NotEqual(t, rands[i], rands[j], "generateRandString caused collision")
					}
				}
			}
		})
	}
}
