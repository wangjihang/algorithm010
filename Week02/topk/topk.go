package topk

import (
	"container/heap"
)

// bitmap
func topKFrequent2(nums []int, k int) []int {
	m := make(map[int]int)
	list := make([][]int, len(nums)+1)
	for _, v := range nums {
		m[v]++
	}
	for k, v := range m {
		list[v] = append(list[v], k)
	}
	var res []int
	for i := len(nums); i >= 1 && len(res) < k; i-- {
		res = append(res, list[i]...)
	}
	return res
}

// 最小的距离
func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	h := &Heap{}
	for _, v := range nums {
		m[v]++
	}
	heap.Init(h)
	for v, c := range m {
		heap.Push(h, &Node{v, c})
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	res := make([]int, 0, k)
	for h.Len() > 0 {
		res = append(res, heap.Pop(h).(*Node).Val)
	}
	return res
}

type Node struct {
	Val   int
	Count int
}

// 最小堆
type Heap []*Node

func (h *Heap) Less(i, j int) bool {
	return (*h)[i].Count < (*h)[j].Count
}

func (h *Heap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *Heap) Len() int {
	return len(*h)
}

func (h *Heap) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

func (h *Heap) Push(v interface{}) {
	*h = append(*h, v.(*Node))
}
