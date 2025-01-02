package main

import "fmt"

func main(){
	arr := []int{5, 2, 8, 1, 9, 3}
	fmt.Println(selectionSort(arr))
}

func selectionSort(arr []int) []int {
	for i := 0; i < len(arr) - 1; i++ {
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
	return arr
}