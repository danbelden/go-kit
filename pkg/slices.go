package pkg

// StringSliceUnique returns the unique strings in a slice of strings
func StringSliceUnique(slice []string) []string {
	var uniqueSlice []string
	uniqueMap := make(map[string]bool)
	for _, sliceString := range slice {
		if _, ok := uniqueMap[sliceString]; ok {
			continue
		}
		uniqueSlice = append(uniqueSlice, sliceString)
		uniqueMap[sliceString] = true
	}
	return uniqueSlice
}
