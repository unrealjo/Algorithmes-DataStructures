package selectionsort

func swap(n1, n2 int, arr []int) {
	temp := arr[n1]
	arr[n1] = arr[n2]
	arr[n2] = temp
}

func SelectionSort(arr []int) []int {
	min := 0
	for i := 0; i < len(arr); i++ {
		min = i
		for j := i + 1; j < len(arr); j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		swap(i, min, arr)
	}
	return arr
}
