package Zuo

import (
	"fmt"
	"math/rand"
)

// Partition逻辑
// 0. 初始小于等于区为-1
// 1. [i] <= num：当前数和小于等于区的下一个数交换，小于等于区++，当前数++
// 2. [i] > num：当前数++

// PartitionEx逻辑
// 小于在左 等于在中 大于在右边
// 0. 初始 小于区 = -1 大于区 = len(arr)
// 1. [i] == num i++
// 2. [i] < num, 交换小于区右一个和[i] i++ 小于区++
// 3. [i] > num, 交换大于区左一个和[i] ××i不变×× 大于区--

// 以arr[r]划分，返回左边界与右边界
func NeitherlandsFlag(arr []int, l, r int) (newL, newR int) {
	if l > r {
		return -1, -1
	}
	if l == r {
		return l, r
	}
	newL = l - 1
	newR = r + 1
	p := l
	nPartition := arr[r]
	for p <= r {
		if arr[p] < nPartition {
			arr[p], arr[newL+1] = arr[newL+1], arr[p]
			newL++
			p++
		} else if arr[p] > nPartition {
			arr[p], arr[newR-1] = arr[newR-1], arr[p]
			newR--
		} else {
			p++
		}
	}
	return
}

func QuickSort(arr []int) {
	if len(arr) < 2 {
		return
	}
	quickSortSimple(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func quickSortSimple(arr []int, l, r int) {
	if l >= r {
		return
	}
	randNum := l + rand.Intn(r-l)
	arr[randNum], arr[r] = arr[r], arr[randNum]
	fmt.Println("randNum", randNum)
	fmt.Println("curArr", arr)
	m := quickSortPartition(arr, l, r)
	quickSortSimple(arr, l, m-1)
	quickSortSimple(arr, m+1, r)
	return
}

func quickSortPartition(arr []int, l, r int) int {
	less := l - 1
	more := r
	index := l
	for index < more {
		if arr[index] < arr[r] {
			arr[index], arr[less+1] = arr[less+1], arr[index]
			index++
			less++
		} else if arr[index] > arr[r] {
			arr[index], arr[more-1] = arr[more-1], arr[index]
			more--
		} else {
			index++
		}
	}
	arr[r], arr[more] = arr[more], arr[r]
	fmt.Println(arr, l, r)
	return less + 1
}

/***************Final Version***************/
func QuickSortEx(arr []int) {
	if nil == arr || len(arr) < 2 {
		return
	}
	fmt.Println("Initial arr", arr)
	quickSortExSim(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func quickSortExSim(arr []int, l, r int) {
	if l >= r {
		return
	}
	randIndex := l + rand.Intn(r-l)
	arr[randIndex], arr[r] = arr[r], arr[randIndex]
	fmt.Println("randNum", randIndex)
	fmt.Println("curArr", arr)
	newL, newR := partitionEx(arr, l, r)
	quickSortExSim(arr, l, newL-1)
	quickSortExSim(arr, newR+1, r)
	return
}

func partitionEx(arr []int, l, r int) (newL, newR int) {
	less := l - 1
	more := r
	index := l
	for index < more {
		if arr[index] > arr[r] {
			arr[index], arr[more-1] = arr[more-1], arr[index]
			more--
		} else if arr[index] < arr[r] {
			arr[index], arr[less+1] = arr[less+1], arr[index]
			less++
			index++
		} else {
			index++
		}
	}
	arr[more], arr[r] = arr[r], arr[more]
	fmt.Println("arr", arr, "l", l, "r", r)
	return less + 1, more
}
