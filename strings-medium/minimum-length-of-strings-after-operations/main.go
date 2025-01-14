package main

import (
	"fmt"
)

func main() {
	s := "waoji"
	fmt.Println(minimumLength(s))
}

func minimumLength(s string) int {
	n := len(s)

	if n < 2 {
		return n
	}
	count := [26]int{}
	for _, ch := range s {
		c := ch - 'a'
		if count[c] == 2 {
			count[c]--
		} else {
			count[c]++
		}
	}
	fmt.Println(count)
	minLen := 0
	for _, val := range count {
		minLen += val
	}

	return minLen
}

func minimumLength2(s string) int {
	n := len(s)

	if n < 2 {
		return n
	}

	right := make([]int, 26)

	for i := 0; i < n; i++ {
		rc := s[i] - 'a'
		if right[rc] == 2 {
			right[rc]--
		} else {
			right[rc]++
		}
	}

	minLen := 0
	for _, val := range right {
		minLen += val
	}

	return minLen
}
