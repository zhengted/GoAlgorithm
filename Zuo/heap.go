package Zuo

import (
	"container/heap"
)

// 数组表示完全二叉树
//  左孩子2×i+1 右孩子 2×(i+1) 父节点 (i-1)/2  [从0开始的]
//  从1开始的也有
//      i << 1      i << 1 | 1     i>>1

// 给定一个数组，数组排序结束后，所有数字距离原来的位置小于0
type IntArrLeeK []int

func (arr IntArrLeeK) Less(i, j int) bool {
	return arr[i] < arr[j]
}

func (arr IntArrLeeK) Len() int {
	return len(arr)
}

func (arr IntArrLeeK) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func (arr *IntArrLeeK) Push(val interface{}) {
	*arr = append(*arr, val.(int))
}

func (arr *IntArrLeeK) Pop() interface{} {
	old := *arr
	n := len(old)
	x := old[n-1]
	*arr = old[0 : n-1]
	return x
}

func SortedArrayDistanceLessK(arr []int, k int) {
	temp := &IntArrLeeK{}
	heap.Init(temp)
	for i := 0; i < k && i < len(arr); i++ {
		heap.Push(temp, arr[i])
	}
	index := 0
	for j := k; j < len(arr); j++ {
		heap.Push(temp, arr[j])
		arr[index] = heap.Pop(temp).(int)
		index++
	}

	for temp.Len() > 0 {
		arr[index] = heap.Pop(temp).(int)
		index++
	}

}

// 前K大问题
type TopKArray []int

func (arr TopKArray) Less(i, j int) bool {
	return arr[i] > arr[j]
}

func (arr TopKArray) Len() int {
	return len(arr)
}

func (arr TopKArray) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func (arr *TopKArray) Push(val interface{}) {
	*arr = append(*arr, val.(int))
}

func (arr *TopKArray) Pop() interface{} {
	old := *arr
	n := len(old)
	x := old[n-1]
	*arr = old[0 : n-1]
	return x
}

func topKNum(nums []int, k int) []int {
	temp := &TopKArray{}
	heap.Init(temp)
	for i := 0; i < len(nums); i++ {
		heap.Push(temp, nums[i])
	}
	res := []int{}
	for j := 0; j < k; j++ {
		num := heap.Pop(temp).(int)
		res = append(res, num)
	}
	return res
}

type Item struct {
	Index    int
	Val      int
	Frequent int
}

type ItemArray []*Item

func (arr ItemArray) Len() int {
	return len(arr)
}

func (arr ItemArray) Less(i, j int) bool {
	return arr[i].Frequent > arr[j].Frequent
}

func (arr ItemArray) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
	arr[j].Index = i
	arr[i].Index = j
}

func (arr *ItemArray) Push(val interface{}) {
	item := val.(*Item)
	item.Index = len(*arr)
	*arr = append(*arr, item)
}

func (arr *ItemArray) Pop() interface{} {
	old := *arr
	n := len(old)
	x := old[n-1]
	x.Index = -1
	old[n-1] = nil
	old = old[0 : n-1]
	*arr = old
	return x
}

func topKFrequent(nums []int, k int) []int {
	return nil
}
