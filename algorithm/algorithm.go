package algorithm

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	list := make([][]int, len(nums)+1)
	for _, v := range nums {
		m[v]++
	}
	for k, v := range m {
		list[v] = append(list[v], k)
	}
	res := make([]int, 0, k)
	for i := len(nums); i > 0 && len(res) < k; i-- {
		res = append(res, list[i]...)
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
	return (*h)[i].Count > (*h)[j].Count
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
