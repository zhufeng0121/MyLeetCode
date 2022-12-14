package SwordToOffer

func getLeastNumbers(arr []int, k int) []int {
	if len(arr) <= k {
		return arr
	}

}

func quicksort(arr []int) []int {
	
}

func partition(arr []int, start, end int, k int) {
	i, j := start, end
	// 取最左元素作为 "哨兵"的索引位置
	pivot := arr[start]

	for i < j {
		// 从右向左，查找首个小于基准数的元素
		for i < j && arr[j] >= pivot {
			j--
		}
		// 从左向右，查找首个大于基准数的元素
		for i < j && arr[i] <= pivot {
			i++
		}

		arr[i], arr[j] = arr[j], arr[i]
	}

	//交换基准数
	arr[start], arr[i] = arr[i], arr[start]

	if i > k {
		partition(arr)
	}

}
