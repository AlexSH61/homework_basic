package counter

import "errors"

func CountWords(preparedString []string) (map[string]int, error) {
	wordCount := make(map[string]int)
	if len(preparedString) == 0 {
		return nil, errors.New("incorrect data")
	}
	for _, word := range preparedString {
		wordCount[word]++
	}

	return wordCount, nil
}
