package utils

func Has[T comparable](s []T, x T) bool {
	for _, e := range s {
		if e == x {
			return true
		}
	}
	return false
}