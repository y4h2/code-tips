package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSegmentTree(t *testing.T) {
	assert := assert.New(t)
	nums := []int{-2, 0, 3, -5, 2, -1}
	st := NewSegmentTree(nums)

	fmt.Println(st.tree)

	result, err := st.Query(0, 2)
	assert.NoError(err)
	assert.Equal(1, result)
	result, err = st.Query(2, 5)
	assert.NoError(err)
	assert.Equal(-1, result)
	result, err = st.Query(0, 5)
	assert.NoError(err)
	assert.Equal(-3, result)
}
