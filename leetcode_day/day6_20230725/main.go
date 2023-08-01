package main

import (
	"container/heap"
	"fmt"
)

//2208. 将数组和减半的最少操作次数
func main() {
	nums := []int{5, 19, 8, 1}
	fmt.Println(halveArray(nums))
}

func halveArray(nums []int) int {
	var pq = &PriorityQueue{}
	var rst int
	var t1, t2 float64
	//制造float数组并排序
	for _, num := range nums {
		t1 += float64(num)
		heap.Push(pq, float64(num))
	}
	for t2 < t1/2 {
		ele := heap.Pop(pq).(float64)
		t2 += ele / 2
		heap.Push(pq, ele/2)
		rst++
	}
	return rst
}

//实现最大堆和最小堆排序
type PriorityQueue []float64

func (pq PriorityQueue) Len() int {
	return len(pq)
}

//">"代表升序，"<"代表降序排序
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i] > pq[j]
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	*pq = append(*pq, x.(float64))
}

func (pq *PriorityQueue) Pop() any {
	old, n := *pq, len(*pq)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}
