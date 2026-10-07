// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

// Package randutil provides primitives for generating random values
package randutil

import (
	crand "crypto/rand"
	mrand "math/rand/v2" // used for non-crypto unique ID and random port selection
	"sync"
)

// MathRandomGenerator is a random generator for non-crypto usage.
type MathRandomGenerator interface {
	// Intn returns random integer within [0:n).
	Intn(n int) int

	// Uint32 returns random 32-bit unsigned integer.
	Uint32() uint32

	// Uint64 returns random 64-bit unsigned integer.
	Uint64() uint64

	// GenerateString returns ranom string using given set of runes.
	// It can be used for generating unique ID to avoid name collision.
	//
	// Caution: DO NOT use this for cryptographic usage.
	GenerateString(n int, runes string) string
}

type mathRandomGenerator struct {
	r  *mrand.Rand
	mu sync.Mutex
}

// NewMathRandomGenerator creates a new mathematical random generator.
// It terminates the process if it cannot obtain a seed from crypto/rand.
func NewMathRandomGenerator() MathRandomGenerator {
	var seed [32]byte
	// crypto/rand.Read fills the buffer or terminates if its source fails. It uses operating
	// system APIs that are documented to never return an error on all but legacy Linux systems.
	// Note to maintainers: if you need to fall back to timestamp-based seeding for legacy systems,
	// put that code behind a build tag to make default builds secure.
	_, _ = crand.Read(seed[:])

	return newMathRandomGenerator(seed)
}

func newMathRandomGenerator(seed [32]byte) *mathRandomGenerator {
	return &mathRandomGenerator{r: mrand.New(mrand.NewChaCha8(seed))} //nolint:gosec // ChaCha8 is seeded by crypto/rand; gosec flags math/rand/v2 broadly.
}

func (g *mathRandomGenerator) Intn(n int) int {
	g.mu.Lock()
	v := g.r.IntN(n)
	g.mu.Unlock()

	return v
}

func (g *mathRandomGenerator) Uint32() uint32 {
	g.mu.Lock()
	v := g.r.Uint32()
	g.mu.Unlock()

	return v
}

func (g *mathRandomGenerator) Uint64() uint64 {
	g.mu.Lock()
	v := g.r.Uint64()
	g.mu.Unlock()

	return v
}

func (g *mathRandomGenerator) GenerateString(n int, runes string) string {
	letters := []rune(runes)
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[g.Intn(len(letters))]
	}

	return string(b)
}
