package linked_list

import "fmt"

type Node struct {
	data int
	next *Node
}
type LinkedList struct {
	head *Node
}

func NewLinkedList() LinkedList {
	return LinkedList{head: nil}
}

func (list LinkedList) IsEmpty() bool {
	return list.head == nil
}
func (list *LinkedList) Append(value int) {
	tmp := &Node{data: value, next: nil}

	if list.head == nil {
		list.head = tmp
		return
	}

	current := list.head

	for current.next != nil {
		current = current.next
	}

	current.next = tmp
}
func (list *LinkedList) Prepend(value int) {
	tmp := &Node{data: value, next: nil}

	if list.head == nil {
		list.head = tmp
		return
	}
	tmp.next = list.head
	list.head = tmp
}
func (list LinkedList) Length() int {
	if list.head == nil {
		return 0
	}

	var count int = 0
	current := list.head

	for current != nil {
		current = current.next
		count++
	}

	return count

}
func (list LinkedList) Get(index int) int {
	if list.head == nil {
		return 0
	}

	current := list.head

	for current != nil && index != 0 {
		current = current.next
		index--
	}

	if index != 0 {
		return 0
	}

	return current.data
}

func (list LinkedList) Search(value int) int {
	if list.head == nil {
		return -1
	}

	current := list.head
	var index int = 0

	for current != nil {
		if current.data == value {
			return index
		}
		current = current.next
		index++
	}

	return -1
}

func (list LinkedList) ToArray() []int {
	if list.head == nil {
		return []int{}
	}
	arr := make([]int, list.Length())
	i := 0
	current := list.head

	for current != nil {
		arr[i] = current.data
		i++
		current = current.next
	}
	return arr
}

func (list LinkedList) Print() {
	if list.head == nil {
		fmt.Println("[]")
	}

	current := list.head

	fmt.Printf("[ ")
	for current != nil {
		fmt.Printf("%d => ", current.data)
		current = current.next
	}
	fmt.Println("nil ]")
}
func (list *LinkedList) Delete(index int) bool {
	if list.head == nil {
		return false
	}
	if index == 0 {
		list.head = list.head.next
		return true
	}
	current := list.head
	for i := 0; i < index-1 && current != nil; i++ {
		current = current.next
	}
	if current.next == nil {
		return false
	}
	current.next = current.next.next
	return true
}
