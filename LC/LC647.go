package LC

import "fmt"

func CountSubstrings(s string) int {
	str := []byte(s)
	ret := 0
	for m := 0; m <= len(str); m++ {
		for n := len(str) - 1; n >= 0 && n >= m; n-- {
			if isSubs(str[m : n+1]) {
				ret++
			}
		}
	}
	return ret
}

func isSubs(str []byte) bool {
	p := 0
	q := len(str) - 1
	for p < q {
		if str[p] != str[q] {
			return false
		}
		p++
		q--
	}
	return true
}

func CountSubstringsDP(s string) int {
	bytes := []byte(s)
	dp := make([][]int, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]int, len(bytes))
	}
	for i := 0; i < len(bytes); i++ {
		dp[i][i] = 1
	}
	for i := 1; i < len(bytes); i++ {
		m := 0
		n := i
		for m < len(bytes) && n < len(bytes) {
			dp[m][n] = maxInt(dp[m][n-1], dp[m+1][n]) + 1
			if isSubs(bytes[m : n+1]) {
				dp[m][n]++
			}
			m++
			n++
		}
	}
	fmt.Println(dp)
	return dp[0][len(bytes)-1]
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
