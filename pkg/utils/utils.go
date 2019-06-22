package utils

// ContextKey is the unique key that represents a context value
type ContextKey string

func (c ContextKey) String() string {
	return "context key " + string(c)
}
