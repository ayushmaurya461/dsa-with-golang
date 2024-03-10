package main

import "fmt"

func main() {
	str := "aab"
	fmt.Println(findPalindrome(str))

}
func findPalindrome(str string) [][]string {
	result := [][]string{}
	palindromes(0, str, []string{}, &result)
	return result
}
func palindromes(ind int, str string, path []string, result *[][]string) {
	if ind == len(str) {
		strCopy := make([]string, len(path))
		copy(strCopy, path)
		*result = append(*result, strCopy)
		return
	}
	for i := ind; i < len(str); i++ {
		if checkPalindrome(str, ind, i) {
			fmt.Println(ind, i-ind+1)
			path = append(path, str[ind:i+1])
			palindromes(i+1, str, path, result)
			path = path[:len(path)-1]
		}
	}
}

func checkPalindrome(str string, start, end int) bool {
	for start <= end {
		if str[start] != str[end] {
			return false
		}
		start = start + 1
		end = end - 1
	}
	return true
}
