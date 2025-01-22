// SPDX-FileCopyrightText: 2023 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package randutil

import (
	"sync"
	"testing"
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
				if err != nil {
					t.Fatal(err)
				}

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

				if len(rands) != maxIterations {
					t.Fatal("Failed to generate randoms")
				}

				for i := 0; i < maxIterations; i++ {
					for j := i + 1; j < maxIterations; j++ {
						if rands[i] == rands[j] {
							t.Fatalf("generateRandString caused collision: %s == %s", rands[i], rands[j])
						}
					}
				}
			}
		})
	}
}
