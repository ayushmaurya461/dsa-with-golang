package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	rotate(arr[:], 3)
	fmt.Println(arr)
}

func rotateLeft(arr []int) {
	n := arr[0]
	for i := 0; i < len(arr)-1; i++ {
		arr[i] = arr[i+1]
	}
	arr[len(arr)-1] = n
}

func rotate(nums []int, k int) {
	n := k % len(nums)
	reverse := func(start, end int) {
		for start < end {
			nums[start], nums[end] = nums[end], nums[start]
			start++
			end--
		}
	}
	fmt.Println(n)

	reverse(0, len(nums)-1)
	reverse(0, n-1)
	reverse(n, len(nums)-1)
}

func rotateLeftByN(arr []int, n int) {
	n = n % len(arr)
	if n == 0 {
		return
	}
	temp := make([]int, n)
	copy(temp, arr[:n])

	for i := n; i < len(arr); i++ {
		arr[i-n] = arr[i]
	}

	for i := 0; i < len(temp); i++ {
		arr[len(arr)-n+i] = temp[i]
	}

}
