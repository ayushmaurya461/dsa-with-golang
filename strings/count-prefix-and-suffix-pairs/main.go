package main

import "fmt"

func main() {
	words := []string{"a", "abb"}
	fmt.Println(countPrefixSuffixPairs(words))
}

func countPrefixSuffixPairs(words []string) int {
	count := 0
	for i := 0; i < len(words); i++ {
		word := words[i]
		for j := i + 1; j < len(words); j++ {
			if len(word) <= len(words[j]) {
				match := checkPrefixSuffix(words[j], word)
				if match {
					count++
				}
			}
		}
	}
	return count
}

func checkPrefixSuffix(str, word string) bool {
	n := len(word)
	strLength := len(str)
	prefix := str[:n]
	suffix := str[strLength-n : strLength]
	return word == prefix && word == suffix
}
