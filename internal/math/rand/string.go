package rand

import "math/rand"

// Based on https://stackoverflow.com/a/31832326/3027614

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

// StringRand is a source of random characters.
type StringRand struct {
	src rand.Source
}

// NewStringRand returns a Rand implementation that generates strings.
func NewStringRand(src rand.Source) *StringRand {
	return &StringRand{src}
}

// Read generates len(p) random bytes (that can be represented as characters)
// and writes them into p. It always returns len(p) and a nil error.
func (r *StringRand) Read(p []byte) (n int, err error) {
	n = len(p)

	// A r.src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, r.src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = r.src.Int63(), letterIdxMax
		}

		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			p[i] = letterBytes[idx]
			i--
		}

		cache >>= letterIdxBits
		remain--
	}

	return
}
