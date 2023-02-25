package quicksort

import (
	"fmt"
	"reflect"
	"testing"
)

type Case struct {
	arr      []int
	expected []int
}

var tests = []Case{
	{[]int{3, 4, 5, 1, 2}, []int{5, 4, 3, 2, 1}},
	{[]int{0, 9, 0, 1}, []int{9, 1, 0, 0}},
}

func TestQuickSort(t *testing.T) {
	for _, test := range tests {
		get := []int{}
		get = append(get, test.arr...)
		QuickSort(get, 0, len(get)-1)
		if !reflect.DeepEqual(get, test.expected) {
			t.Errorf("Expect %v got %d", test.expected, get)
		} else {
			fmt.Printf("Array : %v\t=> Sorted : %v\n", test.arr, get)
		}
	}
}
func BenchmarkQuickSort(b *testing.B) {
	arr := []int{5, 2, 9, 1, 3}
	for i := 0; i < b.N; i++ {
		QuickSort(arr, 0, len(arr)-1)
	}
}
