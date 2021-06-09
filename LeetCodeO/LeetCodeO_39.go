package LeetCodeO

func majorityElement(nums []int) int {
	nCandidate := 0
	nHp := 1
	for i := 1; i < len(nums); i++ {
		if nHp == 0 {
			nCandidate = i
			nHp = 1
		} else {
			if nums[nCandidate] == nums[i] {
				nHp++
			} else {
				nHp--
			}
		}
	}
	return nums[nCandidate]
}
