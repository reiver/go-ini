package iniconv

func MapStrings[T any](strings []string, fn func(string)T) []T {
	var result []T = []T{}

	for _, datum := range strings {
		result = append(result, fn(datum))
	}

	return result
}
