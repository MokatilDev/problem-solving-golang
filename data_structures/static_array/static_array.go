package static_array

import (
	"errors"
	"fmt"
)

type StaticArray struct {
	values []int
	length int
	size   int
}

func (arr *StaticArray) Init(size int) {
	arr.values = make([]int, size)
	arr.size = size
	arr.length = 0
}

func (arr *StaticArray) Len() int {
	return arr.length
}

func (arr *StaticArray) InsertEnd(value int) error {
	if arr.length < arr.size {
		arr.values[arr.length] = value
		arr.length++
		return nil
	}

	return errors.New("Array is Full")
}

func (arr *StaticArray) RemoveEnd() error {
	if arr.length > 0 {
		arr.values[arr.length-1] = 0
		arr.length--
		return nil
	}

	return errors.New("Array is Empty")
}

func (arr *StaticArray) InsertMiddle(value int, index int) error {
	if arr.length >= arr.size {
		return errors.New("Array is Full")
	}

	if index < 0 || index > arr.size {
		return errors.New("Out of range")
	}

	for i := arr.length - 1; i > index; i-- {
		arr.values[i] = arr.values[i-1]
	}

	arr.values[index] = value
	arr.length++

	return nil
}

func (arr *StaticArray) RemoveMiddle(index int) error {
	if arr.length < 0 {
		return errors.New("Array is Empty")
	}

	if index < 0 || index > arr.length {
		return errors.New("Out of range")
	}

	for i := index + 1; i < arr.length-1; i++ {
		arr.values[i-1] = arr.values[i]
	}

	arr.values[arr.length-1] = 0
	arr.length--

	return nil
}

func (arr *StaticArray) PrintArr() {
	for i := range arr.values {
		fmt.Println(arr.values[i])
	}
}
