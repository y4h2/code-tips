package main

import (
	"errors"
	"testing"
)

func BinarySearch(arr []int, target int) (int, error) {
	low, high := 0, len(arr)-1
	for low <= high {
		mid := low + (high-low)>>1
		if target == arr[mid] {
			return mid, nil
		} else if target < arr[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1, errors.New("not found")
}

func TestBinarySearch(t *testing.T) {
	t.Log("Given an array")
	arr := []int{1, 2, 3, 5, 8, 10, 20, 100}

	index, err := BinarySearch(arr, arr[5])
	assert := NewAssert(t)
	assert.NoError(err)
	assert.Equal(5, index)
}

// 找到第一个值等于给定值的元素
func BinarySearchFindFirst(arr []int, target int) (int, error) {
	low, high := 0, len(arr)-1
	for low <= high {
		mid := low + (high-low)>>1
		if arr[mid] > target {
			high = mid - 1
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			if mid == 0 || arr[mid-1] != target {
				return mid, nil
			} else {
				high = mid - 1
			}
		}
	}

	return -1, errors.New("not found")
}

// 查找最后一个值等于给定值的元素
func BinarySearchFindLast(arr []int, target int) (int, error) {
	low, high := 0, len(arr)-1
	for low <= high {
		mid := low + (high-low)>>1
		if arr[mid] < target {
			low = mid + 1
		} else if arr[mid] > target {
			high = mid - 1
		} else {
			if mid == len(arr)-1 || arr[mid+1] != target {
				return mid, nil
			} else {
				low = mid + 1
			}
		}
	}

	return -1, errors.New("not found")
}

// 查找第一个大于等于给定值的元素
func BinarySearchFirstGE(arr []int, target int) (int, error) {
	low, high := 0, len(arr)-1

	for low <= high {
		mid := low + (high-low)>>1
		if arr[mid] >= target {
			if mid == 0 || !(arr[mid-1] >= target) {
				return mid, nil
			} else {
				high = mid - 1
			}
		} else {
			low = mid + 1
		}
	}

	return -1, errors.New("not found")
}

// 查找最后一个小于等于给定值的元素
func BinarySearchLastLE(arr []int, target int) (int, error) {
	low, high := 0, len(arr)-1

	for low <= high {
		mid := low + (high-low)>>1
		if arr[mid] <= target {
			if mid == len(arr)-1 || !(arr[mid+1] <= target) {
				return mid, nil
			} else {
				low = mid + 1
			}
		} else {
			high = mid - 1
		}
	}

	return -1, errors.New("not found")
}

type Assert struct {
	t *testing.T
}

func NewAssert(t *testing.T) *Assert {
	return &Assert{t}
}

func (a *Assert) NoError(err error) {
	if err != nil {
		a.t.Errorf("expect no error but got %s", err.Error())
	}
}

func (a *Assert) Equal(i, j int) {
	if i != j {
		a.t.Errorf("expect %d, got %d", i, j)
	}
}

func matchString(a, b string) (bool, error) {
	return a == b, nil
}

func main() {
	testSuite := []testing.InternalTest{
		{
			Name: "test case 1",
			F:    TestBinarySearch,
		},
	}
	testing.Main(matchString, testSuite, nil, nil)
}
