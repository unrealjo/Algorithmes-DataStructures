package mergesort

func merge(left, right []int) []int {
	var res []int = make([]int, len(left)+len(right))
	var i, j, k, m int
	for i < len(left) && j < len(right) {
		if left[i] >= right[j] {
			res[k] = right[j]
			j++
		} else {
			res[k] = left[i]
			i++
		}
		k++
	}
	for m = i; m < len(left); m++ {
		res[k] = left[m]
		k++
	}
	for m = j; m < len(right); m++ {
		res[k] = right[m]
		k++
	}
	return res

}
func MergeSort(arr []int) []int {
	var size int = len(arr)
	if size <= 1 {
		return arr
	}
	leftArr := MergeSort(arr[:size/2])
	rightArr := MergeSort(arr[size/2:])
	return merge(leftArr, rightArr)
}
