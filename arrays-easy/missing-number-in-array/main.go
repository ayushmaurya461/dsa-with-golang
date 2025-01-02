package main

func main() {
	nums := []int{3, 0, 1}
	println(missingNumber(nums))
}

func missingNumber(nums []int) int {
	xor1 := 0
	xor2 := 0
	for i := 0; i < len(nums); i++ {
		xor1 = xor1 ^ nums[i]
		xor2 = xor2 ^ (i + 1)
	}
	return xor1 ^ xor2
}
