package Zuo

// 提取一个整形数最右侧的1
func GetMostRightOne(n int) int {
	return n & ((^n) + 1)
}

// 一个数组中有两种数出现了奇数次，其他都出现了偶数次
func FindTwoNumberAppearOddTimes(arr []int) (int, int) {
	eor := arr[0]
	for i := 1; i < len(arr); i++ {
		eor ^= arr[i]
	}
	// 异或结束后 eor一定不为0
	//  整个数组分成两部分 一部分某位不为0 一部分某位为0
	rightOne := GetMostRightOne(eor) // 提取出最右侧1
	onlyOne := 0                     //eor'
	for i := 0; i < len(arr); i++ {
		if arr[i]&rightOne != 0 {
			onlyOne ^= arr[i]
		}
	}
	return onlyOne, eor ^ onlyOne
}

func BitCount(n int) int {
	count := 0
	for n != 0 {
		rightOne := GetMostRightOne(n)
		count++
		n ^= rightOne
	}
	return count
}
