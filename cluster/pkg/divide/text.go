package divide

import (
	"strings"
)

func divideTextIntoSentences(text string) []string {
	sentences := strings.Split(text, ".")
	var result []string

	for _, sentence := range sentences {
		sentence = strings.TrimSpace(sentence)
		if sentence != "" {
			result = append(result, sentence)
		}
	}

	return result
}

