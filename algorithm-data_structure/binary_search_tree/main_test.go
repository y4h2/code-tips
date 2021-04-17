package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
			5
		/   \
	 3     6
	/ \     \
 2   4     8
*/

func TestBST(t *testing.T) {
	assert := assert.New(t)
	nums := []int{5, 3, 6, 8, 4, 2}
	bst := NewBST()
	for _, num := range nums {
		bst.Add(num)
	}

	sort.Ints(nums)
	assert.Equal(nums, bst.InOrder())
}

func TestBSTPreOrderIterator(t *testing.T) {
	assert := assert.New(t)
	nums := []int{5, 3, 6, 8, 4, 2}
	bst := NewBST()
	for _, num := range nums {
		bst.Add(num)
	}

	assert.Equal([]int{5, 3, 2, 4, 6, 8}, bst.PreOrderIteration())
}

func TestBSTLevelOrder(t *testing.T) {
	assert := assert.New(t)
	nums := []int{5, 3, 6, 8, 4, 2}
	bst := NewBST()
	for _, num := range nums {
		bst.Add(num)
	}

	assert.Equal([]int{5, 3, 6, 2, 4, 8}, bst.LevelOrder())
}

func TestBSTRemoveMin(t *testing.T) {
	assert := assert.New(t)
	nums := []int{5, 3, 6, 8, 4, 2}
	bst := NewBST()
	for _, num := range nums {
		bst.Add(num)
	}

	bst.RemoveMin()

	assert.Equal([]int{5, 3, 6, 4, 8}, bst.LevelOrder())
}

func TestBSTRemoveMax(t *testing.T) {
	assert := assert.New(t)
	nums := []int{5, 3, 6, 8, 4, 2}
	bst := NewBST()
	for _, num := range nums {
		bst.Add(num)
	}

	bst.RemoveMax()

	assert.Equal([]int{5, 3, 6, 2, 4}, bst.LevelOrder())
}
