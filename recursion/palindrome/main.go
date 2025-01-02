package main

import "fmt"

func main() {
	str := "AABCAA"
	fmt.Println(checkPalindrome(str, len(str)))
	fmt.Println(checkPalindromeUsingRecursion(str, 0, len(str)))

}
func checkPalindrome(str string, length int) bool {
	for index := range str {
		if index >= length/2 {
			return true
		}
		if str[index] != str[length-index-1] {
			return false
		}
	}
	return false
}

func checkPalindromeUsingRecursion(str string, index, length int) bool {
	if index >= length/2 {
		return true
	}
	if str[index] == str[length-index-1] {
		return checkPalindromeUsingRecursion(str, index+1, length)
	}
	return false
}
