package main

import (
	"GoAlgorithm/LC"
	"fmt"
)

func main() {
	fmt.Println(LC.CountSubstringsDP("aaa"))
}

func testPartition() {
	temp := []int{1, 2, 3, 4, 5, 7, 9, 10, 12, 3, 3, 4, 5}
	x := 3
	less := -1
	more := len(temp)
	i := 0
	for i < more {
		fmt.Println(i, temp)
		if temp[i] < x {
			temp[i], temp[less+1] = temp[less+1], temp[i]
			less += 1
			i += 1
		} else if temp[i] > x {
			temp[i], temp[more-1] = temp[more-1], temp[i]
			more -= 1
		} else {
			i++
		}
	}
	fmt.Println(temp)
}
