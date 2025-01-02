package main

import "fmt"

func main() {
	arr := []int{1, 5, 2, 9, 6, 8, 3, 10}

	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func quickSort(arr []int, low, high int) {
	if low < high {
		pivotIndex := sort(arr, low, high)
		quickSort(arr, low, pivotIndex-1)
		quickSort(arr, pivotIndex+1, high)
	}
}

func sort(arr []int, low, high int) int {
	pivot := arr[low]
	i := low
	j := high
	for i <= j {
		fmt.Println(1)
		for arr[i] <= pivot && i <= high {
			i++
		}
		for arr[j] > pivot && j >= low {
			j--
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[low], arr[j] = arr[j], arr[low]
	return j
}
