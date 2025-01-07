package main

func main() {
	nums := []int{10, 4, -8, 7}
	nums2 := []int{2, 3, 1, 0}
	println(waysToSplitArray(nums))
	println(waysToSplitArray(nums2))
}

func waysToSplitArray(nums []int) int {
	ways := 0
	totalSum := 0
	leftSum := 0
	for i := 0; i < len(nums); i++ {
		totalSum += nums[i]
	}

	for i := 0; i < len(nums)-1; i++ {
		leftSum += nums[i]
		totalSum -= nums[i]
		if leftSum >= totalSum {
			ways++
		}
	}

	return ways
}
