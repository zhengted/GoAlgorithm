package main

import (
	"GoAlgorithm/LeetCodeO"
	"GoAlgorithm/Zuo"
	"fmt"
)

func main() {
	fmt.Println(LeetCodeO.TranslateNum(12258))
	nums := []int{7, 6, 2, 9, 4, 6, 3, 7, 2, 5, 1}
	Zuo.DeliverNum(nums)
	fmt.Println(nums)
}
