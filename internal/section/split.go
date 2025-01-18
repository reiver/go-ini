package inisection

func split(name string) []string {
	if "" == name {
		return []string{}
	}

	var result []string

	var part []rune

	var escaped bool = false

	loop: for _, r := range name {
		if !escaped {
			switch r {
			case '\\':
				escaped = true
	/////////////////////// CONTINUE
				continue loop
			case '.':
				result = append(result, string(part))
				part = nil
	/////////////////////// CONTINUE
				continue loop
			}
		}
		escaped = false
		part = append(part, r)
	}

	result = append(result, string(part))

	return result
}
