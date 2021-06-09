package LeetCodeO

import "strconv"

func TranslateNum(num int) int {
	strNum := strconv.Itoa(num)
	bytes := []byte(strNum)
	dp := make([]int, len(strNum)+1)
	dp[0] = 1
	for i := 0; i < len(strNum); i++ {
		k := i + 1
		if k <= 1 {
			dp[k] = 1
			continue
		}

		if ((bytes[i-1]-'0')*10+(bytes[i]-'0') >= 0 && (bytes[i-1]-'0')*10+(bytes[i]-'0') < 10) ||
			((bytes[i-1]-'0')*10+(bytes[i]-'0') > 25 && (bytes[i-1]-'0')*10+(bytes[i]-'0') <= 99) {
			dp[k] = dp[k-1]
		} else {
			dp[k] = dp[k-1] + dp[k-2]
		}
	}
	return dp[len(bytes)]
}
