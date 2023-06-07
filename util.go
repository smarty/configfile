package configfile

func Coalesce[T comparable](values ...T) (zero T) {
	for _, item := range values {
		if item != zero {
			return item
		}
	}
	return zero
}
