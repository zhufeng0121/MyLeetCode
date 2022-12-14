package Hot100

import "sort"

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

func topKFrequentII(nums []int, k int) []int {
	return nil
}
