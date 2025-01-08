package main

import "fmt"

func main() {
	words := []string{"leetcode", "et", "code"}
	fmt.Println(stringMatching(words))
}

func stringMatching(words []string) []string {
	result := []string{}
	seen := map[string]bool{}

	for i := 0; i < len(words); i++ {
		word := words[i]
		for j := 0; j < len(words); j++ {
			if i != j {
				matchIndex := bruteForceStringMatch(word, words[j])
				println(word, matchIndex, words[j])
				if matchIndex != -1 {
					if !seen[words[j]] {
						result = append(result, words[j])
						seen[words[j]] = true
					}
				}
			}
		}
	}
	return result
}

func bruteForceStringMatch(str, pattern string) int {
	for i := 0; i <= len(str)-len(pattern); i++ {
		match := true
		for j := 0; j < len(pattern); j++ {
			if pattern[j] != str[i+j] {
				match = false
				break
			}
		}
		if match {
			return i
		}
	}
	return -1
}
