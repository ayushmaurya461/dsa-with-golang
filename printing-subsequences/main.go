package main

import "fmt"

func main() {
	arr := []int{3, 1, 2}

	seq(0, &arr, []int{})
}

func seq(i int, arr *[]int, sq []int) {

	if i >= len(*arr) {
		fmt.Println(sq)
		return
	}
	sq = append(sq, (*arr)[i])
	seq(i+1, arr, sq)
	println("--------------------------")
	sq = sq[:len(sq)-1]
	seq(i+1, arr, sq)
}
