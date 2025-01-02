package main

import "fmt"

func main() {
	arr := []int{5, 2, 8, 1, 9, 3}
	fmt.Println(insertionSort(arr))
}

func insertionSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		j := i
		for j > 0 && arr[j-1] > arr[j] {
			arr[j-1], arr[j] = arr[j], arr[j-1]
			j--
		}
	}

	return arr
}
