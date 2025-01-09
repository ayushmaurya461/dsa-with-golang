package main

import "fmt"

func main() {
	arr := []int{2, 7, 11, 15}
	target := 13

	fmt.Println(twoSum(arr, target))
}

func twoSum(nums []int, target int) []int {
	hashMap := map[int]int{}

	for i, num := range nums {
		currentValue := target - num
		value, ok := hashMap[currentValue]
		if ok {
			return []int{value, i}
		}
		hashMap[num] = i
	}

	return nil
}
