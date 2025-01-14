package main

import "fmt"

func main() {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Println(findMajorityElement(nums))
}

func findMajorityElement(nums []int) int {
	hash := map[int]int{}
	n := len(nums) / 2
	for _, num := range nums {
		hash[num]++
		if hash[num] > n {
			return num
		}
	}
	return -1
}
