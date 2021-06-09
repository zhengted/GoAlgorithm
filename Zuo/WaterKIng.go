package Zuo

// 找出占数组一半以上的数
func WaterKing(nums []int) int {
	if len(nums) <= 0 {
		return -1
	}
	nHp := 1
	nCandidate := 0
	for i := 1; i < len(nums); i++ {
		if nHp == 0 {
			nCandidate = i
			nHp = 1
		} else if nums[nCandidate] == nums[i] {
			nHp++
		} else {
			nHp--
		}
	}
	if nHp <= 0 {
		return -1
	}
	nHp = 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == nums[nCandidate] {
			nHp++
		}
	}
	if nHp > (len(nums) >> 1) {
		return nums[nCandidate]
	}
	return -1

}
