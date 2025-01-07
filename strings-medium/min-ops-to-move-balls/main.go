// Minimum Number of Operations to Move All Balls to Each Box

package main

import (
	"fmt"
)

func main() {
	boxes := "001011"
	// boxes := "110"
	fmt.Println(minOperations(boxes))

}

func minOperations(boxes string) []int {
	n := len(boxes)
	ops := make([]int, n)

	moves := 0
	balls := 0
	for i := 0; i < n; i++ {
		ops[i] = moves
		if boxes[i] == '1' {
			balls++
		}
		moves += balls
	}
	moves = 0
	balls = 0
	for i := n - 1; i >= 0; i-- {
		ops[i] += moves
		if boxes[i] == '1' {
			balls++
		}
		moves += balls
	}

	return ops
}

// func minOperations(boxes string) []int {
// 	// n := len(boxes)
// 	// ops := make([]int, n)

// 	// leftMoves := 0
// 	// leftBalls := 0
// 	// for i := 0; i < n; i++ {
// 	// 	ops[i] = leftMoves
// 	// 	if boxes[i] == '1' {
// 	// 		leftBalls++
// 	// 	}
// 	// 	leftMoves += leftBalls

// 	// 	fmt.Println("ball = " + string(boxes[i]) + " , leftMoves = " + fmt.Sprint(leftMoves) + " , leftBalls = " + fmt.Sprint(leftBalls))
// 	// }

// 	// rightMoves := 0
// 	// rightBalls := 0
// 	// for i := n - 1; i >= 0; i-- {
// 	// 	ops[i] += rightMoves
// 	// 	if boxes[i] == '1' {
// 	// 		rightBalls++
// 	// 	}
// 	// 	rightMoves += rightBalls
// 	// }

// 	return ops
// }
