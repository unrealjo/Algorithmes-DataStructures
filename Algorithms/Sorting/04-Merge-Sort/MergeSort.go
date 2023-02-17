package mergesort

func merge(left, right []int) []int {
	res := []int{}
	sizeL := len(left)
	sizeR := len(right)
	for i := 0; i < sizeL && i < sizeR; i++ {
		if left[i] > right[i] {
			res = append(res, right[i], left[i])
		} else {
			res = append(res, left[i], right[i])
		}
	}
	if sizeL != sizeR {
		if left[sizeL-1] > right[sizeR-1] {
			res[len(res)-1] = right[sizeR-1]
			res = append(res, left[sizeL-1])
		} else {
			res = append(res, right[sizeR-1])
		}
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
