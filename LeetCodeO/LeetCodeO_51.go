package LeetCodeO

func reversePairs(nums []int) int {
	var mergeSort func(nStart, nEnd int, nums []int) (cnt int)
	mergeSort = func(nStart, nEnd int, nums []int) (cnt int) {
		cnt = 0
		if nStart >= nEnd {
			return
		}
		nMid := (nEnd-nStart)/2 + nStart

		cnt = mergeSort(nStart, nMid, nums) + mergeSort(nMid+1, nEnd, nums)
		tmp := []int{}
		i := nStart
		j := nMid + 1
		for i <= nMid && j <= nEnd {
			if nums[i] <= nums[j] {
				tmp = append(tmp, nums[i])
				cnt += j - (nMid + 1)
				i++
			} else {
				tmp = append(tmp, nums[j])
				j++
			}
		}

		for ; i <= nMid; i++ {
			tmp = append(tmp, nums[i])
			cnt += nEnd - (nMid + 1) + 1
		}

		for ; j <= nEnd; j++ {
			tmp = append(tmp, nums[j])
		}

		for i := nStart; i <= nEnd; i++ {
			nums[i] = tmp[i-nStart]
		}

		return cnt
	}
	return mergeSort(0, len(nums)-1, nums)
}
