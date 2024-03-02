package main

import "fmt"

func main() {

	fmt.Println(sumUsingFunction(4))
}

func sumUsingFunction(num int64) int64 {
	if num == 0 {
		return 1
	}
	return num * sumUsingFunction(num-1)
}
