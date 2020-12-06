package strings

import str "strings"

// UniqueSlice returns the unique strings in a slice of strings
func UniqueSlice(slice []string) []string {
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

// SearchWord searches for a word within a string
func SearchWord(text string, word string) bool {
	if text == "" || word == "" {
		return false
	}
	words := str.Split(text, " ")
	for _, wordsWord := range words {
		if wordsWord == word {
			return true
		}
	}
	return false
}
