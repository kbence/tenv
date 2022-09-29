package tenv

func StringSliceContains(slice []string, item string) bool {
	for _, it := range slice {
		if item == it {
			return true
		}
	}

	return false
}
