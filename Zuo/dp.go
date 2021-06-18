package Zuo

import (
	"fmt"
	"math"
)

//返回菲薄拿起数列第n项
func Fibonacci(n int) int {
	if n <= 2 {
		return 1
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

// 机器人在一行N个位置上走 n >= 2
// 开始机器人在M位置上 1<=m<=n
// 机器人在1位置上只能往左走到2 机器人在n位置上只能往右走到n-1
// 规定机器人必须走K步 最终能来到P位置的方法有多少种
func RobotWalk(n, m, k, p int) int {
	return walk(n, m, k, p)
}

// n:位置范围 固定参数
// cur:当前位置
// rest：剩余步数
// p:目的位置
func walk(n, cur, rest, p int) int {
	if rest == 0 {
		if cur == p {
			return 1
		}
		return 0
	}
	if cur == 1 {
		return walk(n, 2, rest-1, p)
	}
	if cur == n {
		return walk(n, n-1, rest-1, p)
	}
	return walk(n, cur+1, rest-1, p) + walk(n, cur-1, rest-1, p)
}

func RobotWalkCache(n, m, k, p int) int {
	dp := make([][]int, n+1)   // 0-N 是cur的范围
	for i := 0; i < n+1; i++ { // 0-k 是剩余步数的范围
		dp[i] = make([]int, k+1)
	}
	// dp表的含义是：当前在的地方剩余多少步 能到终点的方法
	for i := 0; i <= n; i++ {
		for j := 0; j <= k; j++ {
			dp[i][j] = -1
		}
	}

	return walkCache(n, m, k, p, dp)
}

func walkCache(n, cur, rest, p int, dp [][]int) int {
	if dp[cur][rest] != -1 {
		return dp[cur][rest]
	}
	if rest == 0 {
		if cur == p {
			dp[cur][rest] = 1
			return 1
		}
		dp[cur][rest] = 0
		return 0
	}
	if cur == 1 {
		dp[cur][rest] = walkCache(n, 2, rest-1, p, dp)
		return dp[cur][rest]
	}
	if cur == n {
		dp[cur][rest] = walkCache(n, n-1, rest-1, p, dp)
		return dp[cur][rest]
	}
	dp[cur][rest] = walkCache(n, cur+1, rest-1, p, dp) + walkCache(n, cur-1, rest-1, p, dp)
	return dp[cur][rest]
}

func DpBag(w []int, v []int, bag int) int {
	dp := make([][]int, len(w)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, bag+1)
	}
	for index := len(w) - 1; index >= 0; index-- {
		for rest := 0; rest <= bag; rest++ {
			no := dp[index+1][rest]
			yes := -1
			if rest-w[index] >= 0 {
				yes = dp[index+1][rest-w[index]] + v[index]
			}
			dp[index][rest] = maxInt(yes, no)
		}
	}

	return dp[0][bag]
}

func DpNumTran(str string) int {
	dp := make([]int, len(str)+1)
	dp[len(str)] = 1
	for i := len(str) - 1; i >= 0; i-- {
		if str[i] == '0' {
			dp[i] = 0
			continue
		}
		if str[i] == '1' {
			dp[i] = dp[i+1]
			if i+2 <= len(str) {
				dp[i] += dp[i+2]
			}
			continue
		}
		if str[i] == '2' {
			dp[i] = dp[i+1]
			if i+1 <= len(str) && str[i+1] >= '0' && str[i+1] <= '6' {
				dp[i] += dp[i+2]
			}
			continue
		}
		dp[i] = dp[i+1]
	}
	fmt.Println(dp)
	return dp[0]
}

func DpCards(arr []int) int {
	dpF := make([][]int, len(arr))
	dpS := make([][]int, len(arr))
	for i := 0; i < len(arr); i++ {
		dpF[i] = make([]int, len(arr))
		dpS[i] = make([]int, len(arr))
	}

	for i := 0; i < len(arr); i++ {
		dpF[i][i] = arr[i]
		dpS[i][i] = 0
	}

	for i := 1; i < len(arr); i++ {
		L := 0
		R := i
		for L < len(arr) && R < len(arr) {
			// dp[L][R] = ?
			dpF[L][R] = maxInt(arr[L]+dpS[L+1][R], arr[R]+dpS[L][R-1])
			dpS[L][R] = minInt(dpF[L+1][R], dpF[L][R-1])
			L++
			R++
		}
	}

	return maxInt(dpF[0][len(arr)-1], dpS[0][len(arr)-1])
}

// 给定一个正整数数组，数组中的内容代表面值，给定一个整数  求由面值数组凑的该整数的方法数
func FixValue(arr []int, target int) int {
	return fixValue(arr, 0, target)
}

// 可以自由使用arr[index...]所有的面值，组成rest，有所少种方法
func fixValue(arr []int, index, rest int) int {
	if rest < 0 {
		return 0
	}
	if index == len(arr) {
		if rest == 0 {
			return 1
		}
		return 0
	}

	ret := 0
	for nCount := 0; nCount*arr[index] <= rest; nCount++ {
		ret += fixValue(arr, index+1, rest-nCount*arr[index])
	}

	return ret
}

func dpFixValue(arr []int, target int) int {
	dp := make([][]int, len(arr)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, target+1)
	}
	dp[len(arr)][0] = 1

	for index := len(arr) - 1; index >= 0; index-- {
		for rest := 0; rest <= target; rest++ {
			//dp[index][rest] = ?
			// 方法一：
			dp[index][rest] = 0
			for nCount := 0; nCount*arr[index] <= rest; nCount++ {
				dp[index][rest] += dp[index+1][rest-nCount*arr[index]]
			}
			// 方法二：(剪枝优化)
			dp[index][rest] = dp[index+1][rest]
			if rest-arr[index] >= 0 {
				dp[index][rest] += dp[index][rest-arr[index]]
			}
		}
	}

	return dp[0][target]
}

// 两个字符串的最长公共子序列问题
func CommonSeries(a, b string) int {
	return commonSeriesForce(a, b)
}

func commonSeriesForce(restA, restB string) int {
	fmt.Println(restA, restB)
	if len(restA) == 0 || len(restB) == 0 {
		return 0
	}
	bytesA := []byte(restA)
	bytesB := []byte(restB)

	ret := math.MinInt32
	temp := 0
	for _, firstA := range bytesA {
		temp = 0
		for i, ch := range bytesB {
			if firstA == ch {
				temp += commonSeriesForce(restA[1:], restB[i+1:]) + 1
			}
		}
		ret = maxInt(ret, temp)
	}

	return ret
}

// i,j : a从0到i字符串和b从0到j的公共子序列长度
func CommonSeriesDp(strA, strB string) int {
	dp := make([][]int, len(strA))   // 行为strA的长度
	for i := 0; i < len(strA); i++ { // 列为strB的长度
		dp[i] = make([]int, len(strB))
	}
	bytesA := []byte(strA)
	bytesB := []byte(strB)

	// 处理最左上角
	if bytesA[0] == bytesB[0] {
		dp[0][0] = 1
	}

	// 处理第一行
	for i := 1; i < len(bytesB); i++ {
		if bytesA[0] == bytesB[i] {
			dp[0][i] = dp[0][i-1] + 1
		} else {
			dp[0][i] = dp[0][i-1]
		}
	}

	// 处理第一列
	for j := 1; j < len(bytesA); j++ {
		if bytesB[0] == bytesA[j] {
			dp[j][0] = dp[j-1][0] + 1
		} else {
			dp[j][0] = dp[j-1][0]
		}
	}

	// 处理剩余数据
	for i := 1; i < len(bytesA); i++ {
		for j := 1; j < len(bytesB); j++ {
			temp := maxInt(dp[i-1][j], maxInt(dp[i][j-1], dp[i-1][j-1]))
			if bytesB[j] == bytesA[i] {
				dp[i][j] = temp + 1
			} else {
				dp[i][j] = temp
			}
		}
	}
	fmt.Println(dp)
	return dp[len(bytesA)-1][len(bytesB)-1]
}

// 给定一个数组表示每个人喝完咖啡准备刷杯子的时间
// 只有一台洗杯机，一次只能洗一个杯子，时间耗费a
// 也可以自己挥发干净 时间耗费b 挥发可并行
func WashCup(arr []int, a, b int) int {
	return WashCupDp(arr, a, b)
}

// a,b,drinks 固定变量
// index:  drinks[0...index-1]已经干净了，不用你操心了
// washLine: 洗杯机在该时间点才可用
func washCupForce(drinks []int, a, b, index, washLine int) int {
	if index == len(drinks)-1 {
		return minInt(
			maxInt(washLine, drinks[index])+a, // 假设27喝完，洗杯机12能用，那么得到27+3，假设我11喝完 12才能用 的搭配12+3
			drinks[index]+b,                   // 让杯子自行挥发
		) //两者取较小值
	}

	// wash是我当前的咖啡杯洗完的时间
	wash := maxInt(washLine, drinks[index]) + a // 决定洗了
	next1 := washCupForce(drinks, a, b, index+1, wash)
	p1 := maxInt(wash, next1)

	dry := drinks[index] + b // 决定自行挥发了
	next2 := washCupForce(drinks, a, b, index+1, washLine)
	p2 := maxInt(dry, next2)

	return minInt(p1, p2)
}

func WashCupDp(drinks []int, a, b int) int {
	n := len(drinks)
	dp := make([][]int, n-1)
	if a >= b {
		return drinks[n-1] + b
	}
	limit := 0 // 洗杯机什么时候可用
	for i := 0; i < n; i++ {
		limit = maxInt(limit, drinks[i]) + a
	}
	for i := 0; i < n; i++ {
		dp[i] = make([]int, limit+1)
	}

	for washLine := 0; washLine < limit+1; washLine++ {
		dp[n-1][washLine] = minInt(maxInt(washLine, drinks[n-1])+a, drinks[n-1]+b)
	}

	for index := n - 2; index >= 0; index++ {
		for washLine := 0; washLine < limit+1; washLine++ {
			wash := maxInt(washLine, drinks[index]) + a
			p1 := math.MaxInt32
			if wash <= limit {
				p1 = maxInt(wash, dp[index+1][wash])
			}
			p2 := maxInt(drinks[index]+b, dp[index+1][washLine])

			dp[index][washLine] = minInt(p1, p2)
		}
	}

	return dp[0][0]
}
