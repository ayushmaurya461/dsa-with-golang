package main

import (
	"fmt"
)

func main() {
	str := "011101"
	fmt.Println(maximumScore(str))
}

func maximumScore(score string) int {
	strLength := len(score)
	maxScore := 0
	right := 0
	left := 0
	for i := 0; i < strLength; i++ {
		if score[i] == '1' {
			right++
		}
	}

	for i := 1; i < strLength; i++ {
		if score[i-1] == '0' {
			left++
		}
		if score[i-1] == '1' {
			right--
		}

		if left+right > maxScore {
			maxScore = left + right
		}
	}
	return maxScore

}

// BruteForce Solution

// func maximumScore(score string) int {
// 	strLength := len(score)
// 	value := math.MinInt
// 	for i := 1; i < strLength; i++ {
// 		sum := calculate(score[:i], 0) + calculate(score[i:strLength], 1)
// 		if value < sum {
// 			value = sum
// 		}
// 	}
// 	return value
// }

// func calculate(str string, num int) int {
// 	maxScore := 0
// 	for i := 0; i < len(str); i++ {
// 		val, err := strconv.ParseInt(string(str[i]), 10, 64)
// 		if err != nil {
// 			fmt.Println("Error parsing int:", err)
// 			continue
// 		}
// 		if int(val) == num {
// 			maxScore++
// 		}
// 	}
// 	return maxScore
// }
