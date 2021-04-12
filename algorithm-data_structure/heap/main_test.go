package main

import (
	"fmt"
	"testing"
)

func TestHeap(t *testing.T) {
	// items := []int{1, 2, 3, 10, 9, 5}
	// items := []int{1, 2, 1}
	items := []int{2, 1, 1}

	heap := NewHeap()
	for _, item := range items {
		heap.Insert(item)
	}

	n := heap.Size()
	for i := 0; i < n; i++ {
		fmt.Println(heap.Pop())
	}
}

func TestHeapSort(t *testing.T) {
	items := []int{1, 2, 3, 10, 9, 5}

	heap := &Heap{items: items}

	heap.Sort()
	n := heap.Size()
	for i := 0; i < n; i++ {
		fmt.Println(heap.Pop())
	}
}
