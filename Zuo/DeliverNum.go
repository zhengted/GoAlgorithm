package Zuo

// 一个m+n长度的数组，要求奇数放到奇数下标上 偶数放到偶数下标上
func DeliverNum(nums []int) {
	a, b, p := 0, 1, len(nums)-1
	for a < len(nums) && b < len(nums) {
		if nums[p]%2 == 0 {
			nums[a], nums[p] = nums[p], nums[a]
			a += 2
		} else {
			nums[b], nums[p] = nums[p], nums[b]
			b += 2
		}
	}
	return
}
