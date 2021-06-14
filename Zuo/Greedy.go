package Zuo

import "math"

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
//   1. 先分成

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
