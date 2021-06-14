package Zuo

import "fmt"

func MergeSort(arr []int) {
	merge(arr, 0, len(arr)-1)
}

func merge(arr []int, l, r int) {
	if l >= r {
		return
	}
	mid := l + ((r - l) >> 1)
	merge(arr, l, mid)
	merge(arr, mid+1, r)

	tmp := make([]int, r-l+1)
	p := l
	q := mid + 1
	for p <= mid && q <= r {
		if arr[p] <= arr[q] {
			tmp = append(tmp, arr[p])
			p++
		} else {
			tmp = append(tmp, arr[q])
			q++
		}
	}

	for p <= mid {
		tmp = append(tmp, arr[p])
		p++
	}

	for q <= r {
		tmp = append(tmp, arr[q])
		q++
	}

	for i := 0; i < len(tmp); i++ {
		arr[l+i] = tmp[i]
	}

}

func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 非递归实现
func MergeSortEx(arr []int) {
	if arr == nil || len(arr) <= 0 {
		return
	}
	N := len(arr)
	mergeSize := 1
	for mergeSize < N {
		l := 0
		for l < N {
			m := l + mergeSize - 1
			if m >= N {
				break
			}
			r := IntMin(m+mergeSize, N-1)
			// 合并
			l = r + 1
		}
		// 可加可不加，是为了避免下面的乘2操作溢出
		if mergeSize > N/2 {
			break
		}
		mergeSize <<= 1
	}
}

// 小和问题:找出数组中每个数前面比他小的和 算总和
func SmallAdd(arr []int) int {
	return smallAddMerge(arr, 0, len(arr)-1)
}

func smallAddMerge(arr []int, l, r int) int {
	if l == r {
		return 0
	}
	mid := l + ((r - l) >> 1)
	leftAdd := smallAddMerge(arr, l, mid)
	rightAdd := smallAddMerge(arr, mid+1, r)
	selfMerge := smallAddSelfMerge(arr, l, mid, r)

	return leftAdd + rightAdd + selfMerge
}

func smallAddSelfMerge(arr []int, l, mid, r int) int {
	p := l
	q := mid + 1
	res := 0
	tmp := []int{}
	for p <= mid && q <= r {
		// 这里注意，小和问题的小规模问题解是 左边的数比右边的数小时 求出右边的数里有几个比当前的数大
		// 如果相等先合并右边的数字
		if arr[p] < arr[q] {
			res += (r - q + 1) * arr[p]
			fmt.Printf("res += %d*%d\n", r-q+1, arr[p])
			tmp = append(tmp, arr[p])
			p++
		} else {
			tmp = append(tmp, arr[q])
			q++
		}
	}
	for p <= mid {
		tmp = append(tmp, arr[p])
		p++
	}

	for q <= r {
		tmp = append(tmp, arr[q])
		q++
	}
	fmt.Println(tmp)

	for i := 0; i < r-l; i++ {
		arr[l+i] = tmp[i]
	}
	return res
}

// 逆序对问题
func AntiPair(arr []int) int {
	return antiPairSort(arr, 0, len(arr)-1)
}

func antiPairSort(arr []int, l, r int) int {
	if l == r {
		return 0
	}
	mid := l + ((r - l) >> 1)
	return antiPairSort(arr, l, mid) + antiPairSort(arr, mid+1, r) + antiPairMerge(arr, l, mid, r)
}

func antiPairMerge(arr []int, l, mid, r int) int {
	p := l
	q := mid + 1
	tmp := []int{}
	res := 0
	for p <= mid && q <= r {
		// 与小和问题不一样的是 逆序对是当右组的数小于左边的数时 产生逆序对，因此相等时应当左边先入辅助数组
		if arr[q] <= arr[p] {
			res += mid - p + 1
			fmt.Println("res add", mid-p+1)
			tmp = append(tmp, arr[q])
			q++
		} else {
			tmp = append(tmp, arr[p])
			p++
		}
	}

	for p <= mid {
		tmp = append(tmp, arr[p])
		p++
	}

	for q <= r {
		tmp = append(tmp, arr[q])
		q++
	}

	for i := 0; i < r-l+1; i++ {
		arr[l+i] = tmp[i]
	}
	fmt.Println("tmp:", tmp)

	return res
}
