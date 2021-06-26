package LC

import (
	"fmt"
)

func maxLength(arr []string) int {
	ret := 0
	// 子过程定义：当前位置为之前（0...index-1）已经做好决定了，得到的curRes为
	var f func(index int, curRes string)
	f = func(index int, curRes string) {
		if len(curRes) > ret {
			ret = len(curRes)
		}

		if index == len(arr) {
			return
		}

		f(index+1, curRes)

		// 自检
		bytesId := []byte(arr[index])
		for i := 0; i < len(bytesId); i++ {
			for j := i + 1; j < len(bytesId); j++ {
				if bytesId[i] == bytesId[j] {
					return
				}
			}
		}

		// 判断能不能加 不能加返回
		bytesIndex := []byte(arr[index])
		bytesCur := []byte(curRes)

		bCanAdd := true
		for i := 0; i < len(bytesIndex); i++ {
			for j := 0; j < len(bytesCur); j++ {
				if bytesCur[j] == bytesIndex[i] {
					bCanAdd = false
					break
				}
			}
			if !bCanAdd {
				break
			}
		}

		// 可以加
		if bCanAdd {
			newStr := curRes + arr[index]
			f(index+1, newStr)
		}

	}
	f(0, "")
	return ret
}

func MaxLength() {
	fmt.Println(maxLength([]string{"zog", "nvwsuikgndmfexxgjtkb", "nxko"}))
}

func maxLengthDp(arr []string) int {
	if len(arr) == 1 {
		if checkSelf(arr[0]) {
			return len(arr[0])
		} else {
			return 0
		}
	}

	dp := make([][]int, len(arr))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(arr))
	}
	// dp表的含义 [i,j]范围上所得到的最大长度，斜对角线上的值作为边界，自检不过的为0,自检过了的为自身长度
	for i := 0; i < len(dp); i++ {
		if checkSelf(arr[i]) {
			dp[i][i] = len(arr[i])
		} else {
			dp[i][i] = 0
		}
	}

	for i := 1; i < len(arr); i++ {
		p := 0 // 行
		q := i
		for p < len(arr) && q < len(arr) {

			p++
			q++
		}
	}

	return dp[0][len(arr)-1]
}

func checkSelf(str string) bool {
	for i := 0; i < len(str); i++ {
		for j := i + 1; j < len(str); j++ {
			if str[i] == str[j] {
				return false
			}
		}
	}
	return true
}
