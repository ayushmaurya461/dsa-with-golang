package main

import "fmt"

func main() {
	arr := []int{5, 2, 8, 1, 9, 3}
	n := 8
	fmt.Println(linearSearch(arr, n))
}

func linearSearch(arr []int, n int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == n {
			return i
		}
	}
	return -1
}
