package ini

func mapKeys[T any](m map[string]T) []string {
	if len(m) <= 0 {
		return []string{}
	}

	var keys []string
	{
		for key, _ := range m {
			keys = append(keys, key)
		}

		SortKeys(keys)
	}

	return keys
}
