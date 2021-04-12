package main

import (
	"errors"
)

type Heap struct {
	items []int
}

func NewHeap() *Heap {
	return &Heap{
		items: []int{},
	}
}

// i, 2 * i + 1, 2 * i + 2

func (this *Heap) Insert(item int) {
	this.items = append(this.items, item)

	for i := len(this.items) - 1; i >= 1 && !this.Less(i, (i-1)/2); i = (i - 1) / 2 {
		this.Swap(i, (i-1)/2)
	}
}

func (this *Heap) Less(i, j int) bool {
	return this.items[i] < this.items[j]
}

func (this *Heap) Swap(i, j int) {
	this.items[i], this.items[j] = this.items[j], this.items[i]
}

func (this *Heap) Pop() (int, error) {
	n := len(this.items)
	if n == 0 {
		return -1, errors.New("invalid")
	}
	top := this.items[0]
	this.items[0] = this.items[n-1]
	this.items = this.items[:n-1]

	this.Heapify(0)
	return top, nil
}

func (this *Heap) Heapify(index int) {
	i := index
	n := len(this.items)
	for {
		maxPos := i
		if i*2+1 < n && this.Less(i, i*2+1) {
			maxPos = i*2 + 1
		}
		if i*2+2 < n && this.Less(maxPos, i*2+2) {
			maxPos = i*2 + 2
		}
		if maxPos == i {
			break
		}

		this.Swap(i, maxPos)
		i = maxPos
	}
}

func (this *Heap) Size() int {
	return len(this.items)
}

func (this *Heap) Sort() {
	n := len(this.items) - 1

	for n >= 0 {
		this.Swap(0, n)
		n--
		this.Heapify(0)
	}
}
