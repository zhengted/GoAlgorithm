package Zuo

import "math/rand"

// 对数器
func GenerateRandomIntArray(maxSize, maxValue int) []int {
	res := []int{}
	for i := 0; i < maxSize; i++ {
		res = append(res, rand.Intn(maxValue))
	}
	return res
}
