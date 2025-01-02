package main

import "fmt"

func main() {
	printNumber(1, 5)
}

func printNumber(i, n int64) {
	fmt.Println(i)
	if n == i {
		return
	}
	printNumber(i+1, n)
}
