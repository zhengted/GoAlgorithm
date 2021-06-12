package Zuo

import "fmt"

type CycleQueue struct {
	queue []int
	limit int
	size  int
	pushi int
	pulli int
}

func NewCycleQueue(limit int) *CycleQueue {
	return &CycleQueue{
		queue: []int{},
		limit: limit,
		size:  0,
		pushi: 0,
		pulli: 0,
	}
}

func (q *CycleQueue) Push(x int) {
	if q.size == q.limit {
		fmt.Println("the queue is full")
		return
	}
	q.size++
	q.queue[q.pushi] = x
	q.pushi = q.nextIndex(q.pushi)
}

func (q *CycleQueue) Pop() int {
	if q.size == 0 {
		panic("queue is empty")
	}
	q.size--
	ans := q.queue[q.pulli]
	q.pulli = q.nextIndex(q.pulli)
	return ans
}

func (q *CycleQueue) nextIndex(curIndex int) int {
	if curIndex >= q.limit-1 {
		return 0
	}
	return curIndex + 1
}
