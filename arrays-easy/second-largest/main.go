package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{5, 2, 32, 8, 1, 9, 3, 12}
	fmt.Println(secondLargest(arr))
}

func secondLargest(arr []int) int {
	largest := arr[0]
	sLargest := math.MinInt32
	for i := 1; i < len(arr); i++ {
		if arr[i] > largest {
			sLargest = largest
			largest = arr[i]
		} else if arr[i] > sLargest && arr[i] != sLargest {
			sLargest = arr[i]
		}
	}
	return sLargest
}
