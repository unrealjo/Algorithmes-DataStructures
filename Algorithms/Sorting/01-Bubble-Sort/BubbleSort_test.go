package bubblesort

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
	{[]int{5, 2, 9, 1, 5}, []int{1, 2, 5, 5, 9}},
	{[]int{100, 20, 30, 40, 10}, []int{10, 20, 30, 40, 100}},
}

func TestBubbleSort(t *testing.T) {
	for i, test := range tests {
		if !reflect.DeepEqual(BubbleSort(test.arr), test.expected) {
			t.Error("Unexpected result in case ", i)
		} else {
			fmt.Printf("Array : %v => Sorted : %v\n", test.arr, test.expected)
		}
	}
}
func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BubbleSort([]int{5, 2, 9, 1, 5})
	}
}
