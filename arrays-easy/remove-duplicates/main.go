package main

import "fmt"

func main() {
	arr := []int{1, 1, 3, 3, 5, 6, 6, 6, 8, 8}
	fmt.Println(removeDuplicates(arr))
}

func removeDuplicates(arr []int) []int {
	var result []int
	for i := 0; i < len(arr); i++ {
		if i == 0 || arr[i] != arr[i-1] {
			result = append(result, arr[i])
		}
	}
	return result
}
