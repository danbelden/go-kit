package int64s

// UniqueSlice returns the unique int64s in a slice of int64s
func UniqueSlice(slice []int64) []int64 {
	var uniqueSlice []int64
	uniqueMap := make(map[int64]bool)
	for _, sliceInt := range slice {
		if _, ok := uniqueMap[sliceInt]; ok {
			continue
		}
		uniqueSlice = append(uniqueSlice, sliceInt)
		uniqueMap[sliceInt] = true
	}
	return uniqueSlice
}
