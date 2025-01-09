package main

import (
	"fmt"
	"math"
)

func main() {
	str := "hello"
	fmt.Println(scoreOfString(str))
}

func scoreOfString(str string) int {
	sum := 0
	for i := 0; i < len(str)-1; i++ {
		sum += int(math.Abs(float64(int(str[i]) - int(str[i+1]))))
	}
	return sum
}
