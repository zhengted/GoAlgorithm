package Zuo

import (
	"fmt"
	"math"
	"strconv"
)

// 暴力递归：尝试
// 1.把问题转化为规模缩小了的同类问题的子问题
// 2.有明确不需要继续进行递归的条件 base case
// 3.有当得到了子问题的结果之后的决策过错
// 4.不记录每一个子问题的解

// Hanoi:汉诺塔问题
// 左边N个圆盘移到最右边
// 1. 1～N-1移到中间
// 2. N从左边移到右边
// 3. 1～N-1从中间移到右边
func Hanoi(n int) {
	hanoi(n, "left", "right", "middle")
}

func hanoi(n int, from, to, other string) {
	if n == 1 {
		fmt.Println("Move 1 from " + from + " to " + to)
		return
	}
	hanoi(n-1, from, other, to)                                            // 1.
	fmt.Println("Move " + strconv.Itoa(n) + " from " + from + " to " + to) // 2.
	hanoi(n-1, other, to, from)                                            // 3.
}

// ReverseStack: 不使用额外空间逆序一个栈
func ReverseStack(arr []int) {
	if len(arr) == 0 {
		return
	}
	i := GetStackBottom(&arr)
	ReverseStack(arr)
	arr = append(arr, i)
}

// GetStackBottom:获取栈底元素，并将栈底元素上面的元素往下压一层
func GetStackBottom(arr *[]int) int {
	temp := (*arr)[len((*arr))-1]
	(*arr) = (*arr)[:len((*arr))-1]
	if len((*arr)) == 0 {
		return temp
	} else {
		last := GetStackBottom(arr)
		*arr = append(*arr, temp)
		return last
	}
}

// PrintAllSUbsquence:打印字符串的所有子序列
func PrintAllSubsquence(arr string) {
	bytes := []byte(arr)
	path := []byte{}
	printAllSubsquence(bytes, 0, path)
}

func printAllSubsquence(arr []byte, index int, path []byte) {
	if len(arr) == index {
		// 如果需要不重复字面值就增加一个set参数保证不重复
		fmt.Println(path)
		return
	}
	no := path
	yes := append(path, arr[index])
	printAllSubsquence(arr, index+1, yes)
	printAllSubsquence(arr, index+1, no)
}

// Permutation:交换思想
func Permutation(str string) int {
	res := 0
	permuation(str, 0, &res)
	return res
}

func permuation(str string, index int, res *int) {
	if index == len(str) {
		fmt.Println(str)
		*res++
		return
	}
	bytes := []byte(str)
	for i := index; i < len(bytes); i++ {
		bytes[i], bytes[index] = bytes[index], bytes[i]
		newStr := string(bytes)
		permuation(newStr, index+1, res)
		bytes[i], bytes[index] = bytes[index], bytes[i] // 切片递归两层之后就会被静态化 所以这边先换回来 保证结果符合预期
	}
}

// 剪枝版本
func PermutationWithCut(str string) int {
	res := 0
	permuteWithCut(str, 0, &res)
	return res
}

func permuteWithCut(str string, index int, res *int) {
	if len(str) == index {
		*res++
		return
	}
	visit := [26]bool{false}
	bytes := []byte(str)
	for i := index; i < len(bytes); i++ {
		if !visit[bytes[i]-'a'] {
			visit[bytes[i]-'a'] = true
			bytes[i], bytes[index] = bytes[index], bytes[i]
			permuteWithCut(string(bytes), index+1, res)
			bytes[i], bytes[index] = bytes[index], bytes[i]
		}
	}
}

// 从左往右尝试模型
// 1对应A 2对应B 。。。 111-->AAA或者KA或者AK
// 给定一个只含数字字符串 求转换结果数量
// 动态规划，如果当前字符和他的下一位字符组成的数字是[11,26] 结果+2 否则+1
func NumberTrans(str string) int {
	return numberTrans(str, 0)
}

// 两个分支：一次拿一个数还是拿两个数
// [0...index] 做完决策了 [index+1...]还没决定
func numberTrans(str string, index int) int {
	if index == len(str) {
		// 两种理解
		// 1.没有字符时转换为空字符串
		// 2.获得一种结果
		return 1
	}
	if str[index] == '0' {
		// 单独一个0作为开头无法转换
		return 0
	}
	// 如果当前数字 == 1 且不是最后一位 有两种选择
	// 如果当前数字 == 2 要看下一位是不是大于等于0 且小于等于6的 如果是 有两种  如果不是只有一种
	// 如果当前数字 [3,9] 只有一种选择
	if str[index] == '1' {
		res := numberTrans(str, index+1)
		if index+1 < len(str) {
			res += numberTrans(str, index+2) // 表示让i+2及之后的自由作选择  [...,i,i+1]做完决定了 [i+2...]进递归
		}
		return res
	}
	if str[index] == '2' {
		res := numberTrans(str, index+1)
		if str[index+1] >= '0' && str[index+1] <= '6' && index+1 < len(str) {
			res += numberTrans(str, index+2)
		}
		return res
	}
	// 3..9
	return numberTrans(str, index+1)
}

// 从左往右的尝试模型2
// 在不超过bag的情况下能得到的最多价值是多少
func Bag(weights, values []int, bag int) int {
	return calMaxValue(weights, values, bag, 0, 0, 0)
}

// index:表示当前决策第几个货物
func calMaxValue(weights, values []int, bag, index, curWeight, curVal int) int {
	if curWeight > bag { // 这个判断要放前面 更重要
		return math.MinInt32
	}

	if index == len(weights) {
		return curVal
	}

	maxVal := math.MinInt32
	for i := index; i < len(weights); i++ {
		yes := calMaxValue(weights, values, bag, index+1, curWeight+weights[index], curVal+values[index])
		no := calMaxValue(weights, values, bag, index+1, curWeight, curVal)
		maxVal = maxInt(maxVal, maxInt(yes, no))
	}
	return maxVal
}

// 用剩余空间做
func calMaxValueEx(w, v []int, index, rest int) int {
	if rest < 0 {
		return -1 // 无效方案
	}
	if index == len(w) {
		return 0
	}
	no := calMaxValueEx(w, v, index+1, rest)
	yes := -1
	yesEx := calMaxValueEx(w, v, index+1, rest-w[index]) // 考虑选择了index的货物后是否存在有效方案
	if yesEx != -1 {
		yes = v[index] + yesEx
	}
	return maxInt(no, yes)
}

// 范围上的尝试模型
// 给定一个正整数数组arr A和B依次拿走一张牌 A先拿B后拿，每次只能拿走最左或最右，返回最后获胜者的分数
//  定义一个过程 返回先手在ARR里 L R范围内的最大分数
func firstHand(arr []int, l, r int) int {
	if l == r {
		// 只剩一张牌
		return arr[l]
	}
	// 如果拿走左侧的，则自己的分数是左侧加上下一轮的后手分
	return maxInt(arr[l]+secondHand(arr, l+1, r), arr[r]+secondHand(arr, l, r-1))
}

// 在当前步后手玩家只能在先手挑选的不利情况里计算
func secondHand(arr []int, l, r int) int {
	if l == r {
		return 0 // 只剩一张牌还是后手就没有拿的了
	}
	// 如果先手拿走arr[l] 那么后手能拿到的是firstHand(arr,l+1,r)
	// 如果先手拿走arr[r] 那么后手能拿到的是firstHand(arr,l,r-1)
	// 但是这个过程是对手（当前轮次的先手）决定的，对手一定是选择对自己最不利的
	return minInt(firstHand(arr, l+1, r), firstHand(arr, l, r-1))
}

func GetCardWinner(cardNum []int) int {
	return maxInt(firstHand(cardNum, 0, len(cardNum)-1), secondHand(cardNum, 0, len(cardNum)-1))
}

// n*n大小的棋盘 放n个皇后 要求不同行不同列不同斜线
func NQueen(n int) int {
	if n <= 3 {
		return 0
	}
	record := []int{} // 表示第几行的皇后放在第几列
	for i := 0; i < n; i++ {
		record = append(record, -1)
	}
	// nqueen:record[0...i-1]的皇后已经满足要求了，目前来到第i行
	// 返回摆完所有皇后位之后有几种符合要求的摆法
	var nqueen func(i, n int, rec []int) int
	nqueen = func(i, n int, rec []int) int {
		if i == n {
			// fmt.Println(rec)
			return 1
		}
		// 没有到终止位置
		res := 0
		for j := 0; j < n; j++ {
			// 尝试将第i行的皇后放到该行每个位置
			if isValid(rec, i, j) {
				rec[i] = j
				res += nqueen(i+1, n, rec)
			}
		}
		return res
	}
	return nqueen(0, n, record)
}

func isValid(record []int, x, y int) bool {
	for i := 0; i < x; i++ {
		if y == record[i] {
			return false
		}
		if AbsInt(record[i]-y) == AbsInt(x-i) {
			return false
		}
	}
	return true
}

func AbsInt(a int) int {
	if a < 0 {
		return 0 - a
	}
	return a
}

// 位运算解决八皇后暴力问题
func NQueenEx(n int) int {
	rowLimit := (1 << n) - 1 // 行限制固定变量
	if n == 32 {
		rowLimit = -1
	}
	colLimit := 0   // 列限制
	leftLimit := 0  // 左斜线限制
	rightLimit := 0 // 右斜线限制
	return nQueenEx(rowLimit, colLimit, leftLimit, rightLimit)
}
func nQueenEx(rowLimit, colLimit, leftLimit, rightLimit int) int {
	if colLimit == rowLimit {
		return 1
	}
	//colLimit | leftLimit |rightLimit 总限制
	// ^(colLimit | leftLimit |rightLimit) 去掉干扰的0
	// rowLimit & ^(colLimit | leftLimit |rightLimit) 得到当前能放皇后的1位置
	pos := rowLimit & (^(colLimit | leftLimit | rightLimit)) // 所有可以放皇后的位置 都是1
	mostRightOne := 0
	res := 0
	for pos != 0 {
		mostRightOne = pos & (^pos + 1)
		pos = pos ^ mostRightOne // 也可以用减符号
		res += nQueenEx(rowLimit,
			colLimit|mostRightOne,
			(leftLimit|mostRightOne)<<1,
			(rightLimit|mostRightOne)>>1)
	}
	return res
}
