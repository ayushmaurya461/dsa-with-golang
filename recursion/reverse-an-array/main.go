package main

import "fmt"

func main() {
	arr := []int{1, 2, 8, 4, 5, 6}
	reverse(0, arr, len(arr))
	fmt.Println(arr)
}

func swap(i, j int, arr []int) {
	fmt.Println(i, j)
	tem := arr[i]
	arr[i] = arr[j]
	arr[j] = tem
}

func reverse(i int, arr []int, len int) {
	if i >= len/2 {
		return
	}
	swap(i, len-i-1, arr)
	reverse(i+1, arr, len)
}
