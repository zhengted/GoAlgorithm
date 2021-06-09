package LeetCodeO

import (
	"fmt"
	"math"
)

func SpiralOrder(matrix [][]int) []int {
	directions := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}

	flag := make([][]bool, len(matrix))
	for i := 0; i < len(matrix); i++ {
		flag[i] = make([]bool, len(matrix[0]))
	}
	res := []int{}
	if len(matrix) <= 0 {
		return res
	}
	nBigCycle := int(math.Ceil(((float64(len(matrix[0])) + 0.5) / 2)))
	fmt.Println("nBigCycle", nBigCycle)
	for i := 0; i < nBigCycle; i++ {
		corX := i
		corY := i
		var nSmallCycle int
		for j := 0; j < len(directions); j++ {
			if j%2 == 0 {
				nSmallCycle = len(matrix[0]) - 2*i - 1
			} else {
				nSmallCycle = len(matrix) - 2*i - 1
			}
			fmt.Println("nSmallCycle", nSmallCycle)
			if nSmallCycle <= 0 {
				nSmallCycle = 1
			}
			for k := 0; k < nSmallCycle; k++ {
				fmt.Println("corX", corX, "corY", corY)
				if corX < 0 || corY < 0 || corX >= len(matrix) || corY >= len(matrix[0]) {
					corX = corX + directions[j][0]
					corY = corY + directions[j][1]
					break
				}
				if flag[corX][corY] {
					corX = corX + directions[j][0]
					corY = corY + directions[j][1]
					continue
				}
				flag[corX][corY] = true
				res = append(res, matrix[corX][corY])
				corX = corX + directions[j][0]
				corY = corY + directions[j][1]
			}
		}
	}
	return res
}
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	var (
		rows, columns            = len(matrix), len(matrix[0])
		order                    = make([]int, rows*columns)
		index                    = 0
		left, right, top, bottom = 0, columns - 1, 0, rows - 1
	)

	for left <= right && top <= bottom {
		for column := left; column <= right; column++ {
			order[index] = matrix[top][column]
			index++
		}
		for row := top + 1; row <= bottom; row++ {
			order[index] = matrix[row][right]
			index++
		}
		if left < right && top < bottom {
			for column := right - 1; column > left; column-- {
				order[index] = matrix[bottom][column]
				index++
			}
			for row := bottom; row > top; row-- {
				order[index] = matrix[row][left]
				index++
			}
		}
		left++
		right--
		top++
		bottom--
	}
	return order
}
