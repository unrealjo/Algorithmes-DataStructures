package bubblesort

func swap(n1, n2 int, arr []int) {
	temp := arr[n1]
	arr[n1] = arr[n2]
	arr[n2] = temp
}

func BubbleSort(arr []int) []int {
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < len(arr)-1; i++ {
			if arr[i+1] < arr[i] {
				swap(i, i+1, arr)
				swapped = true
			}
		}
	}
	return arr
}
