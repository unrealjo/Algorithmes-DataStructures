package quicksort

func QuickSort(arr []int, start int, end int) {
	if start >= end {
		return
	}
	var pivot int = end
	left := start
	right := end
	for left < right {
		for arr[left] >= arr[pivot] && left < right {
			left++
		}
		for arr[right] <= arr[pivot] && left < right {
			right--
		}
		arr[left], arr[right] = arr[right], arr[left]
	}
	arr[left], arr[pivot] = arr[pivot], arr[left]
	QuickSort(arr, start, left-1)
	QuickSort(arr, left+1, end)
}
