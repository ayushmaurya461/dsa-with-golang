package main

import "fmt"

func main() {
	arr := []int{2, 4, 1, 6, 3, 8, 0, 9, 5}
	mergeSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func mergeSort(arr []int, low, high int) {
	if low < high {
		mid := (low + high) / 2
		mergeSort(arr, low, mid)
		mergeSort(arr, mid+1, high)
		merge(arr, low, high, mid)
	}
}

func merge(arr []int, low, high, mid int) {
	left := low
	right := mid + 1
	result := []int{}
	for left <= mid && right <= high {
		if (arr)[left] >= (arr)[right] {
			result = append(result, (arr)[right])
			right++
		} else {
			result = append(result, (arr)[left])
			left++
		}
	}
	if left <= mid {
		result = append(result, arr[left:mid+1]...)
	}
	if right <= high {
		result = append(result, arr[right:high+1]...)
	}
	for i, val := range result {
		arr[low+i] = val
	}
}
