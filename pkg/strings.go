package pkg

import "strings"

// StringWordSearch searches for a word within a string
func StringWordSearch(text string, word string) bool {
	if text == "" || word == "" {
		return false
	}
	words := strings.Split(text, " ")
	for _, wordsWord := range words {
		if wordsWord == word {
			return true
		}
	}
	return false
}
