package linked_list

import (
	"reflect"
	"testing"
)

func TestIsEmpty(t *testing.T) {
	list := NewLinkedList()

	if !list.IsEmpty() {
		t.Error("Expected linked list to be empty")
	}
}

func TestLength(t *testing.T) {
	list := NewLinkedList()

	if list.Length() != 0 {
		t.Error("Expected linked list length to be 0")
	}

	list.Append(10)

	if list.Length() != 1 {
		t.Error("Expected linked list length to be 1")
	}
}

func TestAppend(t *testing.T) {
	list := NewLinkedList()

	list.Append(10)

	if list.Length() != 1 {
		t.Error("Expected linked list length to be 1 after appending an element")
	}

	value := list.Get(0)
	if value != 10 {
		t.Errorf("Expected value at position 0 to be 10, got %d", value)
	}
}

func TestSearch(t *testing.T) {
	list := NewLinkedList()

	list.Append(10)
	list.Append(20)

	if list.Search(20) != 1 {
		t.Error("Expected search for value 20 to return 2")
	}
}
func TestToArray(t *testing.T) {
	list := NewLinkedList()

	list.Append(1)
	list.Append(7)
	list.Append(0)
	arr := list.ToArray()

	if !reflect.DeepEqual([]int{1, 7, 0}, arr) {
		t.Errorf("Incorrect array result. Expected: %v, got: %v", []int{1, 7, 0}, arr)
	}
}
func TestDelete(t *testing.T) {
	list := NewLinkedList()

	list.Append(10)
	list.Append(20)
	list.Append(8)
	list.Delete(1)
	arr := list.ToArray()

	if !reflect.DeepEqual([]int{10, 8}, arr) {
		t.Errorf("List after deletion is incorrect. Expected: %v, got: %v", []int{10, 8}, arr)
	}
}
