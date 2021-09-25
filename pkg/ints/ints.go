package ints

// UniqueSlice returns the unique ints in a slice of ints
func UniqueSlice(slice []int) []int {
	var uniqueSlice []int
	uniqueMap := make(map[int]bool)
	for _, sliceInt := range slice {
		if _, ok := uniqueMap[sliceInt]; ok {
			continue
		}
		uniqueSlice = append(uniqueSlice, sliceInt)
		uniqueMap[sliceInt] = true
	}
	return uniqueSlice
}
