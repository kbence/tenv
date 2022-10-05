package tenv

func StringSliceContains(slice []string, item string) bool {
	for _, it := range slice {
		if item == it {
			return true
		}
	}

	return false
}

func StringSliceFilter(slice []string, filter func(string) bool) []string {
	newSlice := []string{}

	for _, it := range slice {
		if filter(it) {
			newSlice = append(newSlice, it)
		}
	}

	return newSlice
}

func FilterStringStartsWith(start string) func(string) bool {
	return func(s string) bool {
		if len(s) < len(start) {
			return false
		}

		return s[:len(start)] == start
	}
}
