package main

import (
	"strings"
)

func countDistinctWords(messages []string) int {
	result := make(map[string]int)
	count := 0
	for _, msg := range messages {
		key := strings.ToLower(msg)
		if _, ok := result[key]; !ok {
			result[key] = 1
		}
		result[key]++
	}

	for _, sum := range result {
		count += result[sum]
	}

	return count
}
