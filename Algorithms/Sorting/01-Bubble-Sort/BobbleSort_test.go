package bubblesort

import (
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

func TestBubbleSort1(t *testing.T) {
	for i, test := range tests {
		if !reflect.DeepEqual(BubbleSort1(test.arr), test.expected) {
			t.Error("Unexpected result in case ", i)
		}
	}
}
func TestBubbleSort2(t *testing.T) {
	for i, test := range tests {
		if !reflect.DeepEqual(BubbleSort2(test.arr), test.expected) {
			t.Error("Unexpected result in case ", i)
		}
	}
}
func BenchmarkBubbleSort1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BubbleSort1([]int{5, 2, 9, 1, 5})
	}
}
func BenchmarkBubbleSort2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BubbleSort2([]int{5, 2, 9, 1, 5})
	}
}
