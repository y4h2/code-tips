package main

import "errors"

type SegmentTree struct {
	data []int
	tree []int
}

func NewSegmentTree(data []int) *SegmentTree {
	tempData := make([]int, len(data))
	copy(tempData, data)

	st := &SegmentTree{
		data: data,
		tree: make([]int, 4*len(data)),
	}

	st.buildSegmentTree(0, 0, len(data)-1)
	return st
}

func (this *SegmentTree) buildSegmentTree(treeIndex, l, r int) {
	if l == r {
		this.tree[treeIndex] = this.data[l]
		return
	}

	leftTreeIndex, rightTreeIndex := this.leftChild(treeIndex), this.rightChild(treeIndex)
	mid := l + (r-l)/2
	this.buildSegmentTree(leftTreeIndex, l, mid)
	this.buildSegmentTree(rightTreeIndex, mid+1, r)

	this.tree[treeIndex] = this.merge(this.tree[leftTreeIndex], this.tree[rightTreeIndex])
}

func (this *SegmentTree) merge(a, b int) int {
	return a + b
}

func (this *SegmentTree) leftChild(i int) int {
	return i*2 + 1
}
func (this *SegmentTree) rightChild(i int) int {
	return i*2 + 2
}

func (this *SegmentTree) Query(l, r int) (int, error) {
	if l < 0 || l >= len(this.data) || r < 0 || r >= len(this.data || l > r) {
		return -1, errors.New("invalid scope")
	}
	return this.query(0, 0, len(this.data)-1, l, r), nil
}

func (this *SegmentTree) query(treeIndex, l, r, queryL, queryR int) int {
	if l == queryL && r == queryR {
		return this.tree[treeIndex]
	}

	mid := l + (r-l)/2
	leftTreeIndex, rightTreeIndex := this.leftChild(treeIndex), this.rightChild(treeIndex)

	if queryL >= mid+1 {
		return this.query(rightTreeIndex, mid+1, r, queryL, queryR)
	} else if queryR <= mid {
		return this.query(leftTreeIndex, l, mid, queryL, queryR)
	}

	return this.merge(
		this.query(leftTreeIndex, l, mid, queryL, mid),
		this.query(rightTreeIndex, mid+1, r, mid+1, queryR),
	)
}

func (this *SegmentTree) Set(index, num int) {

}

func (this *SegmentTree) set(treeIndex, l, r, index, num int) {
	if l == r {
		this.tree[treeIndex] = num
		return
	}

	mid := l + (r-l)/2
	leftTreeIndex, rightTreeIndex := treeIndex*2+1, treeIndex*2+2
	if index >= mid+1 {
		this.set(rightTreeIndex, mid+1, r, index, num)
	} else {
		this.set(leftTreeIndex, l, mid+1, index, num)
	}

	this.tree[treeIndex] = this.merge(this.tree[leftTreeIndex], this.tree[rightTreeIndex])
}
