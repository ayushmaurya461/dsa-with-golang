package main

import "fmt"

func main() {
	str := "abcdeiob"
	pattern := "dei"
	fmt.Println(bruteForce(str, pattern))
}

func bruteForce(str, pattern string) int {
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
