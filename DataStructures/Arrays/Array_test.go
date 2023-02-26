package array

import (
	"reflect"
	"testing"
)

func Test_Array_push(t *testing.T) {
	arr := NewArray(1, 3, 4)
	want := []int{1, 3, 4, 8}
	arr.push(8)
	if !reflect.DeepEqual(arr.elements, want) {
		t.Fatalf("arr.push(8) got %v, want %v ", arr.elements, want)
	}
}

func Test_Array_unshift(t *testing.T) {
	arr := NewArray(1, 3, 4)
	want := []int{8, 1, 3, 4}
	arr.unshift(8)
	if !reflect.DeepEqual(arr.elements, want) {
		t.Fatalf("arr.unshift(8) got %v, want %v ", arr.elements, want)
	}
}
func Test_Array_pop(t *testing.T) {
	tests := []struct {
		in   Array[int]
		want int
	}{
		{in: *NewArray(9, 0, 8, 16, 62), want: 62},
		{in: Array[int]{}, want: 0},
	}
	for _, tt := range tests {
		got := tt.in.pop()
		if got != tt.want {
			t.Fatalf("arr.pop() got %v, want %v ", got, tt.want)
		}
	}
}
func Test_Array_shift(t *testing.T) {
	tests := []struct {
		in   Array[int]
		want int
	}{
		{in: *NewArray(9, 0, 8, 16), want: 9},
		{in: Array[int]{}, want: 0},
	}
	for _, tt := range tests {
		got := tt.in.shift()
		if got != tt.want {
			t.Fatalf("arr.shift() got %v, want %v ", got, tt.want)
		}
	}
}
func Test_Array_splice(t *testing.T) {
	arr := NewArray(5, 1, 9)
	want := []int{5, 18, 20, 9}
	arr.splice(1, 1, 18, 20)
	if !reflect.DeepEqual(arr.elements, want) {
		t.Fatalf("arr.splice(1, 1, 18, 20) got %v, want %v ", arr.elements, want)
	}
}

func Test_Array_slice(t *testing.T) {
	arr := NewArray(9, 0, 8, 16, 62)
	want := []int{8, 16}
	got := arr.slice(2, 4)
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("arr.slice(2,4) got %v, want %v ", got, want)
	}
}
func Test_Array_indexOf(t *testing.T) {
	tests := []struct {
		in     Array[int]
		search int
		want   int
	}{
		{in: *NewArray(9, 0, 8, 16, 62), search: 8, want: 2},
		{in: *NewArray(9, 0, 8, 16, 62), search: 1, want: -1},
	}
	for _, tt := range tests {
		got := tt.in.indexOf(tt.search)
		if got != tt.want {
			t.Fatalf("arr.indexOf(%d) got %v, want %v ", tt.search, got, tt.want)
		}
	}
}
func Test_Array_reverse(t *testing.T) {
	arr := NewArray(5, 1, 9)
	want := []int{9, 1, 5}
	arr.reverse()
	if !reflect.DeepEqual(arr.elements, want) {
		t.Fatalf("arr.reverse() got %v, want %v ", arr.elements, want)
	}
}
func Test_Array_sort(t *testing.T) {
	arr := NewArray(5, 1, 9)
	want := []int{9, 5, 1}
	arr.sort()
	if !reflect.DeepEqual(arr.elements, want) {
		t.Fatalf("arr.sort() got %v, want %v ", arr.elements, want)
	}
}
