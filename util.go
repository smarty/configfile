package configfile

import "cmp"

// Deprecated
func Coalesce[T comparable](values ...T) T {
	return cmp.Or(values...)
}
