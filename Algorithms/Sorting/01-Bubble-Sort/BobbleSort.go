package bubblesort

func swap(n1, n2 int, arr []int) {
	temp := arr[n1]
	arr[n1] = arr[n2]
	arr[n2] = temp
}

// method 1
func BubbleSort1(arr []int) []int {
	size := len(arr)
	for i := 0; i < size; i++ {
		for j := 0; j < size-1; j++ {
			if arr[j] > arr[i] {
				swap(i, j, arr)
			}
		}
	}
	return arr
}

// method 2
func BubbleSort2(arr []int) []int {
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
