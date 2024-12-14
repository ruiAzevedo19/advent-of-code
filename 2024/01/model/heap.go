package model

import (
	"errors"
)

var (
	ErrHeapIsEmpty = errors.New("heap is empty")
)

// MinHeap defines a min-heap.
type MinHeap struct {
	values []int
}

// NewMinHeap returns a new heap.
func NewMinHeap() (minHeap *MinHeap) {
	return &MinHeap{
		values: []int{},
	}
}

// Heapify converts a slice into a min-heap.
func Heapify(s []int) (heap *MinHeap) {
	heap = NewMinHeap()

	for _, value := range s {
		heap.Push(value)
	}

	return heap
}

// Push adds an element to the heap.
func (h *MinHeap) Push(value int) {
	h.values = append(h.values, value)
	h.bubbleUp(len(h.values) - 1)
}

// Pop removes the minimum value of the heap.
func (h *MinHeap) Pop() (value int, err error) {
	if len(h.values) == 0 {
		return 0, ErrHeapIsEmpty
	}

	value = h.values[0]

	h.values[0] = (h.values)[len(h.values)-1]
	h.values = (h.values)[:len(h.values)-1]

	h.bubbleDown(0)

	return value, nil
}

// bubbleUp restores the heap property after adding an element.
func (h *MinHeap) bubbleUp(index int) {
	for i := index; i > 0; {
		parent := parent(i)
		if h.values[i] >= h.values[parent] {
			return
		}

		h.values[i], h.values[parent] = h.values[parent], h.values[i]
		i = parent
	}
}

// bubbleDown restores the heap property after removing an element.
func (h *MinHeap) bubbleDown(i int) {
	for {
		leftChild := leftChild(i)
		rightChild := rightChild(i)

		if leftChild >= len(h.values) {
			return
		}

		var child int
		if rightChild >= len(h.values) || h.values[rightChild] > h.values[leftChild] {
			child = leftChild
		} else {
			child = rightChild
		}

		if h.values[child] < h.values[i] {
			h.values[i], h.values[child] = h.values[child], h.values[i]
		}

		i = child
	}
}

// parent returns the parent index of a node.
func parent(i int) int {
	return (i - 1) / 2
}

// leftChild returns the left child index of a node.
func leftChild(i int) int {
	return 2*i + 1
}

// rightChild returns the right child index of a node.
func rightChild(i int) int {
	return 2 * (i + 1)
}
