package Zuo

import (
	"container/heap"
	"math"
)

// 会议安排问题
// 给定一个meeting数组，求24h内能安排的最多的meeting数
type meeting struct {
	nStart int
	nEnd   int
}

func arrangeMeeting(arr []meeting) int {
	if len(arr) == 0 {
		return 0
	}
	return arrange(arr, 0, 0)
}

// 递归解法
// arr：剩余会议  done：已经安排的回忆数量  timeLine：当前时间点
func arrange(arr []meeting, done, timeLine int) int {
	if len(arr) == 0 {
		return done
	}
	max := done
	for i := 0; i < len(arr); i++ {
		if arr[i].nStart >= timeLine {
			newArr := append(arr[:i], arr[i+1:]...)
			max = maxInt(max, arrange(newArr, done+1, arr[i].nEnd))
		}
	}
	return max
}

// 给定一个字符数组，里面的字符只有可能取值'X'或'.'
// ‘.’表示一户人家可以放灯 且会照亮前后两户人家
// 'x'表示墙面不会被照亮也不能放灯 求最小的放灯数
func minLight(arr []byte) int {
	if len(arr) == 0 {
		return 0
	}
	return minLightProcess(arr, 0, []int{})
}

// arr[index:] 自由选择放灯还是不放灯
// arr[0:index] 已经做好决定了，那些放了灯的位置都存在lights里了
func minLightProcess(arr []byte, index int, lights []int) int {
	if index == len(arr) {
		// 结束的时候
		for i := 0; i < len(arr); i++ {
			if arr[i] != 'X' {
				// 存在没有被照亮的一户
				if !lightContain(lights, i-1) && !lightContain(lights, i) && !lightContain(lights, i+1) {
					return math.MaxInt32
				}
			}
		}
		return len(lights)
	} else {
		no := minLightProcess(arr, index+1, lights)
		yes := math.MaxInt32
		if arr[index] == '.' {
			lights = append(lights, index)
			yes = minLightProcess(arr, index+1, lights)
			lights = lights[:len(lights)-1]
		}
		return minInt(yes, no)
	}
}

// 判断light里是否有index位置上的灯
func lightContain(light []int, index int) bool {
	for _, i := range light {
		if i == index {
			return true
		}
	}
	return false
}

// 金条分割
//  给定数组{10,20,30}表示三个人想要的金条数  金条长度60
//   1. 堆存储
//   2. 每次弹出最小的两个 合并并放入堆中
type GoldenHeap []int

func (h GoldenHeap) Len() int {
	return len(h)
}

func (h GoldenHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h GoldenHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *GoldenHeap) Push(val interface{}) {
	*h = append(*h, val.(int))
}

func (h *GoldenHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	old = old[:n-1]
	*h = old
	return x
}

func SeparateGolden(arr []int) int {
	if len(arr) < 1 {
		return 0
	}
	if len(arr) == 1 {
		return arr[0]
	}
	h := &GoldenHeap{}
	heap.Init(h)
	for _, n := range arr {
		heap.Push(h, n)
	}
	ret := 0
	for len(*h) > 1 {
		n1 := heap.Pop(h).(int)
		n2 := heap.Pop(h).(int)
		n3 := n1 + n2
		ret += n3
		heap.Push(h, n3)
	}
	return ret
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
