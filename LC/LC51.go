package LC

import "fmt"

func solveNQueens(n int) [][]string {
	rowLimit := 1<<n - 1
	ret := [][]string{}
	var f func(rowLimit, colLimit, leftLimit, rightLimit int, curRes []string)
	f = func(rowLimit, colLimit, leftLimit, rightLimit int, curRes []string) {
		if rowLimit == colLimit {
			ret = append(ret, curRes)
			return
		}
		pos := rowLimit & (^(colLimit | leftLimit | rightLimit)) // 所有可以放皇后的位置 都是1
		mostRightOne := 0

		for pos != 0 {
			mostRightOne = pos & (^pos + 1)
			pos = pos ^ mostRightOne // 也可以用减符号
			curBytes := make([]byte, n)
			QPos := getBit(mostRightOne)

			for i := 0; i < n; i++ {
				if i == QPos {
					curBytes[i] = 'Q'
				} else {
					curBytes[i] = '.'
				}
			}

			nextStrArr := append(curRes, string(curBytes))
			f(rowLimit,
				colLimit|mostRightOne,
				(leftLimit|mostRightOne)<<1,
				(rightLimit|mostRightOne)>>1, nextStrArr)
		}
	}
	f(rowLimit, 0, 0, 0, []string{})

	for i := 0; i < len(ret); i++ {
		for j := 0; j < n; j++ {
			fmt.Println(ret[i][j])
		}
		fmt.Println("------------------")
	}

	return ret
}

// 辅助函数 给定一个二进制只有一位是1的数 返回1 所在的位数
func getBit(n int) int {
	ret := 0
	cur := 1
	for n != cur {
		cur <<= 1
		ret++
	}
	return ret
}

func NQueen() {
	solveNQueens(7)
}

func Queen0619(n int) int {
	rowLimit := 1<<n - 1
	colLimit := 0
	leftLimit := 0
	rightLimit := 0
	res := 0
	queen(rowLimit, colLimit, leftLimit, rightLimit, &res)
	return res
}

func queen(rowLimit, colLimit, leftLimit, rightLimit int, res *int) {
	if rowLimit == colLimit {
		// 所有的列都填上了
		*res += 1
		return
	}
	possiblePos := rowLimit & (^(colLimit | leftLimit | rightLimit))
	mostRightOne := 0
	for possiblePos != 0 {
		mostRightOne = possiblePos & (^possiblePos + 1)
		possiblePos ^= mostRightOne
		queen(rowLimit, colLimit|mostRightOne, (leftLimit|mostRightOne)<<1, (rightLimit|mostRightOne)>>1, res)
	}
	return
}
