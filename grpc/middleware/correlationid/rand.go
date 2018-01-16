package correlationid

import (
	"math/rand"
	"sync"
	"time"
)

// Based on https://stackoverflow.com/a/31832326/3027614

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

type randGenerator struct {
	length int
	src    rand.Source
}

// NewRandGenerator returns a unique ID generator of the given length.
func NewRandGenerator(length int) *randGenerator {
	return &randGenerator{
		length: length,
		src:    &lockedSource{src: rand.NewSource(time.Now().UnixNano())},
	}
}

func (g *randGenerator) Generate() string {
	b := make([]byte, g.length)

	// A g.src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := g.length-1, g.src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = g.src.Int63(), letterIdxMax
		}

		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}

		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

// lockedSource allows a random number generator to be used by multiple goroutines concurrently.
// The code originates from math/rand.lockedSource, which is unfortunately not exposed.
type lockedSource struct {
	lk  sync.Mutex
	src rand.Source
}

func (r *lockedSource) Int63() (n int64) {
	r.lk.Lock()
	n = r.src.Int63()
	r.lk.Unlock()
	return
}

func (r *lockedSource) Seed(seed int64) {
	r.lk.Lock()
	r.src.Seed(seed)
	r.lk.Unlock()
}
