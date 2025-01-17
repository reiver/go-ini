package iniconv

func FilterStrings(strings []string, fn func(string)bool) []string {
	var result []string = []string{}

	for _, datum := range strings {
		if !fn(datum) {
	////////////// CONTINUE
			continue
		}
		result = append(result, datum)
	}

	return result
}
