package array

import (
	"log"
	"math"
)

type Array[T int | float64] struct {
	elements []T
	length   int
}

func NewArray[T int | float64](params ...T) *Array[T] {
	e := []T{}
	e = append(e, params...)
	return &Array[T]{
		elements: e,
		length:   len(e),
	}
}

func (arr *Array[T]) push(e T) {
	temp := make([]T, arr.length+1)
	i := 0
	for ; i < arr.length; i++ {
		temp[i] = arr.elements[i]
	}
	temp[i] = e
	arr.elements = temp
	arr.length += 1
}

func (arr *Array[T]) unshift(e T) {
	temp := make([]T, arr.length+1)
	for i := 0; i < arr.length; i++ {
		temp[i+1] = arr.elements[i]
	}
	temp[0] = e
	arr.elements = temp
	arr.length += 1
}

func (arr *Array[T]) pop() T {
	if arr.length < 1 {
		log.Println("The array is empty !")
		return 0
	}
	lastElement := arr.elements[arr.length-1]
	temp := make([]T, arr.length-1)
	for i := 0; i < arr.length-1; i++ {
		temp[i] = arr.elements[i]
	}
	arr.elements = temp
	arr.length -= 1
	return lastElement
}
func (arr *Array[T]) shift() T {
	if arr.length < 1 {
		log.Println("The array is empty !")
		return 0
	}
	firstElement := arr.elements[0]
	temp := make([]T, arr.length-1)
	for i := 1; i < arr.length; i++ {
		temp[i-1] = arr.elements[i]
	}
	arr.elements = temp
	arr.length -= 1
	return firstElement
}

func (arr *Array[T]) splice(start, remove int, elements ...T) {
	if start < 0 || start >= arr.length {
		log.Fatalln("Range out of bounders")
	}
	if remove > arr.length {
		log.Fatalln("No enough elements to be removed")
	}
	temp := make([]T, arr.length-remove+len(elements))
	k, i := 0, 0

	for i = 0; i < start; i++ {
		temp[k] = arr.elements[i]
		k++
	}
	for _, e := range elements {
		temp[k] = e
		k++
	}
	for i = start + remove; i < arr.length; i++ {
		temp[k] = arr.elements[i]
		k++
	}
	arr.elements = temp
	arr.length = k
}

func (arr *Array[T]) slice(start, end int) []T {
	if start < 0 || start >= arr.length {
		log.Fatalln("Invalid start argument")
	}
	if start > end || end > arr.length {
		log.Fatalln("Invalid end argument")
	}
	temp := make([]T, int(math.Abs(float64(start-end))))
	k := 0
	for i := start; i < end; i++ {
		temp[k] = arr.elements[i]
		k++
	}
	return temp
}
func (arr *Array[T]) indexOf(search T) int {
	for i, element := range arr.elements {
		if element == search {
			return i
		}
	}
	return -1
}
func (arr *Array[T]) reverse() {
	temp := make([]T, arr.length)
	k := 0
	for i := arr.length - 1; i >= 0; i-- {
		temp[k] = arr.elements[i]
		k++
	}
	arr.elements = temp
}
func (arr *Array[T]) sort() {
	swapped := true
	ptr := &arr.elements
	for swapped {
		swapped = false
		for i := 0; i < arr.length-1; i++ {
			if (*ptr)[i+1] >= (*ptr)[i] {
				(*ptr)[i+1], (*ptr)[i] = (*ptr)[i], (*ptr)[i+1]
				swapped = true
			}
		}
	}
}
