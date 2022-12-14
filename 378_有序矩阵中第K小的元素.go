package main

import "container/heap"

type pq struct {
	data []interface{}
	len  int
}

func (p *pq) Len() int {
	return p.len
}

func (p *pq) Less(a, b int) bool {
	return p.data[a].(int) > p.data[b].(int)
}

func (p *pq) Swap(a, b int) {
	p.data[a], p.data[b] = p.data[b], p.data[a]
}

func (p *pq) Head() interface{} {
	return p.data[0]
}

func (p *pq) Pop() interface{} {
	p.len--
	return p.data[p.len]
}

func (p *pq) Push(o interface{}) {
	p.data[p.len] = o
	p.len++
}

func kthSmallestV(matrix [][]int, k int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0

	}

	// 大顶堆，堆顶元素是最大的元素
	pq := &pq{data: make([]interface{}, k)}

	heap.Init(pq)

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if pq.Len() < k {
				heap.Push(pq, matrix[i][j])
			} else if matrix[i][j] < pq.Head().(int) {
				heap.Pop(pq)
				heap.Push(pq, matrix[i][j])
			} else {
				// 这一步 剪枝，不再进行判断
				break
			}
		}
	}
	return heap.Pop(pq).(int)
}
