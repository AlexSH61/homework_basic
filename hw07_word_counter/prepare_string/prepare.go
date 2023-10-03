package prepare

import (
	"errors"
	"regexp"
	"strings"
)

func StringPrepare(text string) ([]string, error) {
	if len(text) == 0 {
		return nil, errors.New("incorrect text")
	}
	wordSlice := []string{}
	textLower := strings.ToLower(text)
	strRex := regexp.MustCompile(`[[:punct:]]`)
	cleanString := strRex.ReplaceAllString(textLower, "")
	wordArray := strings.Fields(cleanString)
	wordSlice = append(wordSlice, wordArray...)

	return wordSlice, nil
}
