package insertionsort

func InsertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		var current int = arr[i]
		var j int = i - 1
		for j >= 0 && arr[j] > current {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = current
	}
	return arr
}
