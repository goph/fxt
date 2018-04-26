package context

// Following the pattern seen on https://medium.com/@matryer/context-keys-in-go-5312346a868d

type contextKey string

func (c contextKey) String() string {
	return "fxt context key " + string(c)
}
