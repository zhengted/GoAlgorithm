package main

import (
	"GoAlgorithm/Common"
	"fmt"
)

func main() {
	//fmt.Println(LeetCodeO.TranslateNum(12258))
	nums := []int{7, 6, 2, 9, 4, 6, 3, 7, 2, 5, 1}
	//Zuo.InsertionSort(nums)
	fmt.Println(nums)

	a, b := 5, 5
	Common.Swap(&a, &b)
	fmt.Println(a, b)
}
