package counter

import (
	"regexp"
	"strings"
)

func CountWords(text string) map[string]int {
	textLower := strings.ToLower(text)
	strRex := regexp.MustCompile(`[[:punct:]]`)
	cleanString := strRex.ReplaceAllString(textLower, " ")
	wordString := strings.Join(strings.Fields(cleanString), " ")
	sliceWord := strings.Fields(wordString)
	wordCount := make(map[string]int)
	if len(sliceWord) == 0 {
		return wordCount
	}
	for _, word := range sliceWord {
		wordCount[word]++
	}
	return wordCount
}
