package utils

func Fold[T any](input []T, start T, iterator func(current, next T) T) T {
	current := start
	for _, val := range input {
		current = iterator(current, val)
	}
	return current
}

func Contains[T comparable](input []T, element T) bool {
	for _, lhs := range input {
		if lhs == element {
			return true
		}
	}
	return false
}
