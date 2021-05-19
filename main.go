package main

import (
	"fmt"
)

func FibTopDown(n int) int {
	memo := map[int]int{
		0: 0,
		1: 1,
	}
	// 2 - 10
	if val, exist := memo[n]; exist {
		return val
	} else {
		return FibTopDown(n-1) + FibTopDown(n-2)
	}
	return memo[n]
}

func main() {
	// var money int
	// arrChange := []int{1, 10, 20, 50, 100, 200, 500, 1000, 2000, 5000, 10000}
	// sort.Sort(sort.Reverse(sort.IntSlice(arrChange)))
	// fmt.Scanln(&money)

	// fmt.Println("Kembalian: ")
	fmt.Println(FibTopDown(7))
}
