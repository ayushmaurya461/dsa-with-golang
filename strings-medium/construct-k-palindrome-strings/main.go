package main

import (
	"fmt"
)

func main() {
	str := "leetcode"
	fmt.Println(canConstruct(str, 3))
}

func canConstruct(s string, k int) bool {
	if k > len(s) {
		return false
	}
	charsCount := make([]int, 26)
	for _, c := range s {
		charsCount[c-'a']++
	}
	singles := 0
	for _, count := range charsCount {
		if count%2 != 0 {
			singles++
		}
	}
	return singles < k+1
}

func isPalindrome(str string) bool {
	end := len(str) - 1
	start := 0
	for start < end {
		fmt.Println(str[start], str[end])
		if str[start] != str[end] {
			return false
		}
		start++
		end--
	}
	return true
}
