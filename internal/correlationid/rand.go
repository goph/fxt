package correlationid

import (
	"time"

	"github.com/goph/fxt/internal/math/rand"
)

type randGenerator struct {
	length int
	r      *rand.StringRand
}

// NewRandGenerator returns a unique ID generator of the given length.
func NewRandGenerator(length int) *randGenerator {
	return &randGenerator{
		length: length,
		r:      rand.NewStringRand(rand.NewLockedSource(time.Now().UnixNano())),
	}
}

func (g *randGenerator) Generate() string {
	b := make([]byte, g.length)

	g.r.Read(b)

	return string(b)
}
