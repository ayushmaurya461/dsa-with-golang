package main

import "fmt"

func main() {
	s := "bbcbaba"
	fmt.Println(countPalindromicSubsequence(s))
}

func countPalindromicSubsequence(s string) int {
	count := 0
	pmap := map[string]bool{}
	for i := 0; i < len(s)-2; i++ {
		for j := i + 1; j < len(s)-1; j++ {
			for k := j + 1; k < len(s); k++ {
				if s[i] == s[k] {
					subseq := string(s[i]) + string(s[j]) + string(s[k])
					if !pmap[subseq] {
						pmap[subseq] = true
						count++
					}
				}
			}
		}
	}
	return count
}
