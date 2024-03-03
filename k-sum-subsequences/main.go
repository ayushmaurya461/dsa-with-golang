package main

import "fmt"

func main() {
	arr := []int{2, 4, 1, 3, 6}
	target := 11
	seqSum(0, target, 0, len(arr), &arr, []int{})
	fmt.Println("__________________________________________")
	firstSeqSum(0, target, 0, len(arr), &arr, []int{})
	fmt.Println("__________________________________________")
	fmt.Println(countSubSeq(0, target, 0, len(arr), &arr))
}

func seqSum(i, target, sum, n int, arr *[]int, seq []int) {

	if i >= n {
		if sum == target {
			fmt.Println(seq)
		}
		return
	}

	seq = append(seq, (*arr)[i])
	sum += (*arr)[i]
	seqSum(i+1, target, sum, n, arr, seq)
	seq = seq[:len(seq)-1]
	sum -= (*arr)[i]
	seqSum(i+1, target, sum, n, arr, seq)

}

func firstSeqSum(i, target, sum, n int, arr *[]int, seq []int) bool {

	if i >= n {
		if sum == target {
			fmt.Println(seq)
			return true
		} else {
			return false
		}
	}

	seq = append(seq, (*arr)[i])
	sum += (*arr)[i]
	result := firstSeqSum(i+1, target, sum, n, arr, seq)
	if result {
		return true
	}
	seq = seq[:len(seq)-1]
	sum -= (*arr)[i]
	result = firstSeqSum(i+1, target, sum, n, arr, seq)
	return result
}

func countSubSeq(i, target, sum, n int, arr *[]int) int {

	if i >= n {
		if sum == target {
			return 1
		}
		return 0
	}

	sum += (*arr)[i]
	l := countSubSeq(i+1, target, sum, n, arr)

	sum -= (*arr)[i]
	r := countSubSeq(i+1, target, sum, n, arr)

	return l + r
}
