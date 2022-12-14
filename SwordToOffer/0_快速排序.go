package SwordToOffer

// Quicksort 快速排序
func Quicksort(arr []int, low, high int) {
	if low < high {
		pos := partitionI(arr, low, high)
		Quicksort(arr, low, pos-1)
		Quicksort(arr, pos+1, high)
	}

}

func partitionI(arr []int, low, high int) int {
	pivot := arr[low]

	i, j := low, high

	for i < j {
		for i < j && arr[j] >= pivot {
			j--
		}
		for i < j && arr[i] <= pivot {
			i++
		}
		arr[i], arr[j] = arr[j], arr[i]

	}
	arr[low] = pivot
	return low
}
