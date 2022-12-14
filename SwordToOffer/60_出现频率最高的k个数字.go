package SwordToOffer

import (
	"container/heap"
	"sort"
)

func topKFrequent(nums []int, k int) []int {
	countMap := make(map[int]int)
	arr := make([]int, 0)

	for _, v := range nums {
		countMap[v]++
	}

	for k, _ := range countMap {
		arr = append(arr, k)
	}

	sort.Slice(arr, func(a, b int) bool {
		return countMap[arr[a]] > countMap[arr[b]]
	})
	return arr[:k]

}

// 利用 堆 heap 可以研究一下 golang 里面堆的写法
func topKFrequentI(nums []int, k int) []int {
	occurrences := map[int]int{}
	for _, num := range nums {
		occurrences[num]++
	}
	h := &IHeap{}
	heap.Init(h)
	for key, value := range occurrences {
		heap.Push(h, [2]int{key, value})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	ret := make([]int, k)
	for i := 0; i < k; i++ {
		ret[k-i-1] = heap.Pop(h).([2]int)[0]
	}
	return ret
}

type IHeap [][2]int

func (h IHeap) Len() int           { return len(h) }
func (h IHeap) Less(i, j int) bool { return h[i][1] < h[j][1] }
func (h IHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

func (h *IHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
