package main

import (
	"container/list"
	"errors"
)

type Node struct {
	val   int
	left  *Node
	right *Node
}

func NewNode(val int) *Node {
	return &Node{
		val: val,
	}
}

type BST struct {
	root *Node
	size int
}

func NewBST() *BST {
	return &BST{
		root: nil,
		size: 0,
	}
}

func (this *BST) Add(val int) {
	this.root = this.add(this.root, val)
}

func (this *BST) add(node *Node, val int) *Node {
	if node == nil {
		return NewNode(val)
	}

	if val < node.val {
		node.left = this.add(node.left, val)
	} else {
		node.right = this.add(node.right, val)
	}

	return node
}

func (this *BST) Contains(val int) bool {
	return this.contains(this.root, val)
}

func (this *BST) contains(node *Node, val int) bool {
	if node == nil {
		return false
	}

	if val == node.val {
		return true
	} else if val < node.val {
		return this.contains(node.left, val)
	}

	return this.contains(node.right, val)
}

func (this *BST) PreOrder() []int {
	return this.preOrder(this.root)
}

func (this *BST) preOrder(node *Node) []int {
	if node == nil {
		return []int{}
	}

	result := []int{node.val}
	result = append(result, this.preOrder(node.left)...)
	return append(result, this.preOrder(node.right)...)
}

func (this *BST) InOrder() []int {
	return this.inOrder(this.root)
}

func (this *BST) inOrder(node *Node) []int {
	if node == nil {
		return []int{}
	}

	result := this.inOrder(node.left)
	result = append(result, node.val)
	return append(result, this.inOrder(node.right)...)
}

func (this *BST) ToString() {

}

func (this *BST) PreOrderIteration() []int {
	if this.root == nil {
		return []int{}
	}
	stack := list.New()
	result := []int{}
	stack.PushBack(this.root)
	for stack.Len() > 0 {
		cur := stack.Back().Value.(*Node)
		result = append(result, cur.val)
		stack.Remove(stack.Back())

		if cur.right != nil {
			stack.PushBack(cur.right)
		}

		if cur.left != nil {
			stack.PushBack(cur.left)
		}
	}

	return result
}

func (this *BST) InOrderIteration() []int {
	if this.root == nil {
		return []int{}
	}
	result := []int{}
	stack := list.New()
	cur := this.root
	for cur != nil || stack.Len() > 0 {
		for cur != nil {
			stack.PushBack(cur)
			cur = cur.left
		}

		for stack.Len() > 0 {
			cur = stack.Back().Value.(*Node)
			stack.Remove(stack.Back())
			result = append(result, cur.val)
			cur = cur.right
		}
	}

	return []int{}
}

func (this *BST) LevelOrder() []int {
	queue := list.New()
	queue.PushBack(this.root)
	result := []int{}

	for queue.Len() > 0 {
		cur := queue.Front().Value.(*Node)
		queue.Remove(queue.Front())
		result = append(result, cur.val)

		if cur.left != nil {
			queue.PushBack(cur.left)
		}
		if cur.right != nil {
			queue.PushBack(cur.right)
		}
	}

	return result
}
func (this *BST) Min() (int, error) {
	if this.root == nil {
		return -1, errors.New("no node")
	}

	return this.min(this.root).val, nil
}

func (this *BST) min(node *Node) *Node {
	if node.left == nil {
		return node
	}

	return this.min(node.left)
}

func (this *BST) Max() (int, error) {
	if this.root == nil {
		return -1, errors.New("no node")
	}

	return this.max(this.root).val, nil
}

func (this *BST) max(node *Node) *Node {
	if node.right == nil {
		return node
	}

	return this.max(node.right)
}

func (this *BST) RemoveMin() (int, error) {
	val, err := this.Min()
	if err != nil {
		return val, err
	}
	this.root = this.removeMin(this.root)
	return val, nil
}

// 返回删除了最小值之后的根节点
func (this *BST) removeMin(node *Node) *Node {
	if node.left == nil {
		rightNode := node.right
		node.right = nil
		return rightNode
	}

	node.left = this.removeMin(node.left)
	return node
}

func (this *BST) RemoveMax() (int, error) {
	val, err := this.Max()
	if err != nil {
		return val, err
	}

	this.root = this.removeMax(this.root)

	return val, nil
}

// 返回删除了最大值之后的根节点
func (this *BST) removeMax(node *Node) *Node {
	if node.right == nil {
		leftNode := node.left
		node.left = nil
		return leftNode
	}

	node.right = this.removeMax(node.right)
	return node
}

func (this *BST) Remove(val int) {
	this.root = this.remove(this.root, val)
}

// 删除只有左孩子和只有右孩子的节点较为简单
// 叶子节点同理
// 删除左右都有孩子的节点 d, 找到s = min(d->right)
// s是d的后继
// s->right = delMin(d->right)
func (this *BST) remove(node *Node, val int) *Node {
	if node == nil {
		return nil
	}
	if val < node.val {
		node.left = this.remove(node.left, val)
		return node
	} else if val > node.val {
		node.right = this.remove(node.right, val)
		return node
	}
	// node.val == val
	if node.left == nil {
		rightNode := node.right
		node.right = nil
		return rightNode
	}

	if node.right == nil {
		leftNode := node.left
		node.left = nil
		return leftNode
	}

	successor := this.min(node.right)
	successor.right = this.removeMin(node.right)
	successor.left = node.left

	node.left, node.right = nil, nil
	return successor
}

// func (this *BST) Add(val int) {
// 	if this.root == nil {
// 		this.root = NewNode(val)
// 		this.size++
// 	} else {
// 		this.add(this.root, val)
// 	}
// }

// func (this *BST) add(node *Node, val int) {
// 	if val == node.val {
// 		return
// 	} else if val < node.val && node.left == nil {
// 		node.left = NewNode(val)
// 		this.size++
// 		return
// 	} else if val > node.val && node.right == nil {
// 		node.right = NewNode(val)
// 		this.size++
// 		return
// 	}

// 	if val < node.val {
// 		this.add(node.left, val)
// 	} else {
// 		this.add(node.right, val)
// 	}
// }
