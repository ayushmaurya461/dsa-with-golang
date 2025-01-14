package main

import "fmt"

func main() {
	A := []int{1, 3, 2, 4}
	B := []int{3, 1, 2, 4}

	fmt.Println(findThePrefixCommonArray(A, B))
}

func findThePrefixCommonArray(A []int, B []int) []int {
	n := len(A)
	arr := make([]int, n)
	result := make([]int, n)
	count := 0
	for i := 0; i < n; i++ {
		a := A[i]
		b := B[i]
		arr[a-1]++
		arr[b-1]++
		if arr[a-1] == 2 {
			count++
		}
		if a != b && arr[b-1] == 2 {
			count++
		}
		result[i] = count
	}
	return result
}
