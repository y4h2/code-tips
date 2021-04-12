package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTopK(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([]int{0}, topKFrequent([]int{3, 0, 1, 0, 1, 0}, 1))
	assert.Equal([]int{1, 2}, topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
}
