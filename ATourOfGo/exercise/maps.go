package exercise

import (
	"strings"
)

/*
Implement WordCount.
It should return a map of the counts of each “word” in the string s.
The wc.Test function runs a test suite against the provided function and prints success or failure.

You might find strings.Fields helpful.
 */

func WordCount(s string) map[string]int {
	wordCnt := make(map[string]int)

	for _, w := range strings.Fields(s) {
		wordCnt[w]++
	}
	return wordCnt
}
