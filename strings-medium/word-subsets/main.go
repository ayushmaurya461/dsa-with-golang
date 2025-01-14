package main

import (
	"fmt"
	"math"
)

func main() {
	words1 := []string{"amazon", "apple", "facebook", "google", "leetcode"}
	words2 := []string{"e", "o"}
	fmt.Println(wordSubsets(words1, words2))
}

func wordSubsets(words1 []string, words2 []string) []string {
	required := make(map[rune]int)
	for _, word := range words2 {
		wordCount := getCharCount(word)
		for char, freq := range wordCount {
			required[char] = int(math.Max(float64(required[char]), float64(freq)))
		}
	}

	var result []string
	for _, word := range words1 {
		wordCount := getCharCount(word)
		isUniversal := true
		for char, freq := range required {
			if wordCount[char] < freq {
				isUniversal = false
				break
			}
		}
		if isUniversal {
			result = append(result, word)
		}
	}

	return result
}

func getCharCount(word string) map[rune]int {
	count := make(map[rune]int)
	for _, char := range word {
		count[char]++
	}
	return count
}
