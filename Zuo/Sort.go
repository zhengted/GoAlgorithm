package Zuo

import "GoAlgorithm/Common"

func SelectionSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		min := i
		for j := i; j < len(arr); j++ {
			if arr[j] < min {
				min = j
			}
		}
		arr[min], arr[i] = arr[i], arr[min]
	}
}

func BubbleSort(arr []int) {
	for i := len(arr) - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

// 不熟悉
func InsertionSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i - 1; j >= 0; j-- {
			if arr[j+1] < arr[j] {
				Common.Swap(&(arr[j+1]), &(arr[j]))
			}
		}
	}
}
