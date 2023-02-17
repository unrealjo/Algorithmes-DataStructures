package mergesort

import (
	"reflect"
	"testing"
)

type Case struct {
	arr      []int
	expected []int
}

var tests = []Case{
	{[]int{5, 2, 9, 1, 6}, []int{1, 2, 5, 6, 9}},
	{[]int{100, 20, 30, 40, 10}, []int{10, 20, 30, 40, 100}},
}

func TestMergeSort(t *testing.T) {
	var get []int
	for _, test := range tests {
		get = MergeSort(test.arr)
		if !reflect.DeepEqual(get, test.expected) {
			t.Errorf("Expect %v got %d", test.expected, get)
		}
	}
}
func BenchmarkMergeSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MergeSort([]int{5, 2, 9, 1, 5})
	}
}
