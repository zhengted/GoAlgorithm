package Zuo

// 矩阵处理 宏观调度, 不要被下标所局限

func PrintMatrixByZigZag(matrix [][]int) {
	Ar, Ac, Br, Bc, Endr, Endc := 0, 0, 0, 0, len(matrix)-1, len(matrix[0])-1
	fromUp := false // 表示是否从右上往左下打印
	for Ar != Endr+1 {
		// printLevel(matrix,Ar,Ac,Br,Bc,fromUp)
		if Ac == Endc {
			Ar++
		} else {
			Ac++
		}
		if Br == Endr {
			Bc++
		} else {
			Br++
		}
		fromUp = !fromUp
	}
}

// 转圈打印矩阵

// ××原地××旋转正方形矩阵
