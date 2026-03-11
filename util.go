package configfile

import "cmp"

// Coalesce returns the first non-zero value, or if none were provided, the zero value.
// Deprecated
//
//go:fix inline
func Coalesce[T comparable](values ...T) T {
	return cmp.Or(values...)
}
