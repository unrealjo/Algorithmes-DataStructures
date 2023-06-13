package linked_list

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

func (list *LinkedList) IsEmpty() bool {
	return list.head == nil
}
func (list *LinkedList) Append(value int) {
	tmp := Node{data: value, next: nil}

	if list.head == nil {
		list.head = &tmp
		return
	}

	current := list.head

	for current.next != nil {
		current = current.next
	}

	current.next = &tmp
}
func (list *LinkedList) Length() int {
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
func (list *LinkedList) Get(index int) int {
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
func (list *LinkedList) Search(value int) int {
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
