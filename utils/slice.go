package utils

func Has[T comparable](s []T, x T) bool {
	for _, e := range s {
		if e == x {
			return true
		}
	}
	return false
}

func Spread[T any](s ...[]T) (ret []T) {
	for _, s2 := range s {
		ret = append(ret, s2...)
	}
	return ret
}
