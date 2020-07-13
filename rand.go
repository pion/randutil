package randutil

import (
	crand "crypto/rand"
	"encoding/binary"
	"math/big"
	mrand "math/rand" // used for non-crypto unique ID and random port selection
	"sync"
	"time"
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

// NewMathRandomGenerator creates new mathmatical random generator.
// Random generator is seeded by crypto random.
func NewMathRandomGenerator() MathRandomGenerator {
	var seed int64
	if err := binary.Read(crand.Reader, binary.LittleEndian, &seed); err != nil {
		// crypto/rand is unavailable. Fallback to seed by time.
		seed = time.Now().UnixNano()
	}

	return &mathRandomGenerator{r: mrand.New(mrand.NewSource(seed))}
}

func (g *mathRandomGenerator) Intn(n int) int {
	g.mu.Lock()
	v := g.r.Intn(n)
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

// GenerateCryptoRandomString generates a random string for cryptographic usage.
func GenerateCryptoRandomString(n int, runes string) (string, error) {
	letters := []rune(runes)
	b := make([]rune, n)
	for i := range b {
		v, err := crand.Int(crand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		b[i] = letters[v.Int64()]
	}
	return string(b), nil
}
