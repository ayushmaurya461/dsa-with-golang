package main

import "fmt"

func main() {
	arr := []int{5, 2, 8, 1, 9, 3}
	fmt.Println(bubbleSort(arr))
}

func bubbleSort(arr []int) []int {
	for i := len(arr) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}
