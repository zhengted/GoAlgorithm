package Zuo

// 二分

// 查询一个数是否存在
func BSExist(arr []int, num int) bool {
	L, R := 0, len(arr)-1
	mid := 0
	for L <= R {
		mid = L + ((R - L) >> 1)
		if arr[mid] == num {
			return true
		} else if arr[num] > num {
			R = mid - 1
		} else {
			L = mid + 1
		}
	}
	return arr[L] == num
}

// 在ARR上找满足 >= value的最左位置
func BSNearLeft(arr []int, value int) (ret int) {
	l, r := 0, len(arr)-1
	for l <= r {
		mid := l + ((r - l) >> 1)
		if arr[mid] >= value {
			ret = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return
}

// 局部最小值问题 无序也能二分
func BSAwesome(arr []int) (index int) {
	if len(arr) < 2 {
		return 0
	}
	if arr[0] < arr[1] {
		return 0
	}
	if arr[len(arr)-1] < arr[len(arr)-2] {
		return len(arr) - 1
	}
	l, r := 0, len(arr)-1
	for l <= r {
		mid := l + ((l - r) >> 2)
		if arr[mid] < arr[mid+1] && arr[mid] < arr[mid-1] {
			return mid
		}
		if arr[mid] > arr[mid+1] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l
}
