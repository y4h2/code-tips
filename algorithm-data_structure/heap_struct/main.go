package main

import (
	"errors"
)

type Item struct {
	count int
	num   int
}

type Heap struct {
	items []Item
}

func NewHeap() *Heap {
	return &Heap{
		items: []Item{},
	}
}

func (this *Heap) Push(item Item) {
	this.items = append(this.items, item)

	this.ShiftUp(len(this.items) - 1)
}

func (this *Heap) Pop() (Item, error) {
	n := len(this.items)
	if n == 0 {
		return Item{}, errors.New("not found")
	}

	result := this.items[0]
	this.items[0] = this.items[n-1]
	this.items = this.items[:n-1]
	this.ShiftDown(0)

	return result, nil
}

func (this *Heap) ShiftUp(index int) {
	for i := index; i != (i-1)/2 && !this.Cmp((i-1)/2, i); i = (i - 1) / 2 {
		this.Swap(i, (i-1)/2)
	}
}

func (this *Heap) ShiftDown(i int) {
	n := len(this.items)
	for {
		maxPos, leftPos, rightPos := i, i*2+1, i*2+2
		if leftPos < n && !this.Cmp(i, leftPos) {
			maxPos = leftPos
		}
		if rightPos < n && !this.Cmp(maxPos, rightPos) {
			maxPos = rightPos
		}

		if maxPos == i {
			return
		}

		this.Swap(i, maxPos)
		i = maxPos
	}
}

func (this *Heap) Cmp(i, j int) bool {
	return this.items[i].count > this.items[j].count
}

func (this *Heap) Swap(i, j int) {
	this.items[i], this.items[j] = this.items[j], this.items[i]
}

func topKFrequent(nums []int, k int) []int {
	hash := map[int]int{}
	for _, num := range nums {
		hash[num]++
	}

	heap := NewHeap()
	for k, v := range hash {
		heap.Push(Item{
			count: v,
			num:   k,
		})
	}

	result := make([]int, k)
	for i := 0; i < k; i++ {
		item, _ := heap.Pop()
		result[i] = item.num
	}

	return result
}
