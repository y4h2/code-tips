package main

type UnionFind struct {
	parent []int
}

func NewUnionFind(size int) *UnionFind {
	uf := &UnionFind{parent: make([]int, size)}

}
