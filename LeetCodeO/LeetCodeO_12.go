package LeetCodeO

func exist(board [][]byte, word string) bool {
	bytes := []byte(word)
	if len(bytes) <= 0 {
		return false
	}
	direction := [][]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}
	flag := [][]bool{}
	for i := 0; i < len(board); i++ {
		temp := []bool{}
		for j := 0; j < len(board[0]); j++ {
			temp = append(temp, false)
		}
		flag = append(flag, temp)
	}

	var dfs func(x, y int, wordLeft string) bool
	dfs = func(x, y int, wordLeft string) bool {
		if len(wordLeft) == 0 {
			return true
		}
		flag[x][y] = true
		for i := 0; i < len(direction); i++ {
			newX := x + direction[i][0]
			newY := y + direction[i][1]
			if newX < 0 || newY < 0 || newX > len(board)-1 || newY > len(board[0])-1 {
				continue
			}
			if true == flag[newX][newY] {
				continue
			}
			if board[newX][newY] != wordLeft[0] {
				continue
			}
			if dfs(newX, newY, wordLeft[1:]) {
				flag[x][y] = false
				return true
			}

		}
		flag[x][y] = false
		return false
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == bytes[0] && dfs(i, j, word[1:]) {
				return true
			}
		}
	}

	return false
}
