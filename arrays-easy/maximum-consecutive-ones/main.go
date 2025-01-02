package main

func main() {
	nums := []int{1, 1, 0, 1, 1, 1}
	println(findMaxConsecutiveOnes(nums))
}

func findMaxConsecutiveOnes(nums []int) int {
	max := 0
	count := 0
	for _, num := range nums {
		if num == 1 {
			count++
		} else {
			count = 0
		}
		if count > max {
			max = count
		}
	}
	return max
}
