package rand

import (
	"math/rand"
	"sync"
)

// lockedSource allows a random number generator to be used by multiple goroutines concurrently.
// The code originates from math/rand.lockedSource, which is unfortunately not exposed.
type lockedSource struct {
	lk  sync.Mutex
	src rand.Source
}

// NewLockedSource returns a new pseudo-random Source seeded with the given value.
// The returned source is safe for concurrent usage.
func NewLockedSource(seed int64) rand.Source {
	return &lockedSource{src: rand.NewSource(seed)}
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
