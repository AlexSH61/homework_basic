package counter

import (
	"regexp"
	"strings"
)

var strRex = regexp.MustCompile(`[[:punct:]]`)

func cleanString(text string) string {
	return strRex.ReplaceAllString(text, "")
}

func CountWords(text string) map[string]int {
	textLower := strings.ToLower(text)
	cleanText := cleanString(textLower)
	sliceWord := strings.Fields(cleanText)
	wordCount := make(map[string]int)
	if len(sliceWord) == 0 {
		return wordCount
	}
	for _, word := range sliceWord {
		wordCount[word]++
	}
	return wordCount
}
