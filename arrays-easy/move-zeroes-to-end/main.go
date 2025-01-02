package main

import "fmt"

func main() {
	arr := []int{5, 2, 0, 8, 0, 0, 1, 9, 0, 3}
	moveZeroes(arr)
	fmt.Println((arr))
}

func moveZeroes(arr []int) {
	j := -1
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			j = i
			break
		}
	}
	for i := j; i < len(arr); i++ {
		if arr[i] != 0 {
			arr[i], arr[j] = arr[j], arr[i]
			j++
		}
	}
}
