package main

import "container/list"

// 玩转算法面试6.3
type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{
		val: val,
	}
}

type Command struct {
	s    string // go, print
	node *TreeNode
}

func NewCommand(s string, node *TreeNode) *Command {
	return &Command{
		s:    s,
		node: node,
	}
}

func PreOrderTraverse(root *TreeNode) []int {
	result := []int{}

	stack := list.New()
	stack.PushBack(NewCommand("go", root))

	for stack.Len() > 0 {
		command := stack.Back().Value.(*Command)
		stack.Remove(stack.Back())

		if command.s == "print" {
			result = append(result, command.node.val)
		} else { // command.s == "go"
			if command.node.right != nil {
				stack.PushBack(NewCommand("go", command.node.right))
			}
			if command.node.left != nil {
				stack.PushBack(NewCommand("go", command.node.left))
			}

			stack.PushBack(NewCommand("print", command.node))
		}
	}

	return result
}

func InOrderTraverse(root *TreeNode) []int {
	result := []int{}

	stack := list.New()
	stack.PushBack(NewCommand("go", root))

	for stack.Len() > 0 {
		command := stack.Back().Value.(*Command)
		stack.Remove(stack.Back())

		if command.s == "print" {
			result = append(result, command.node.val)
		} else { // command.s == "go"
			if command.node.right != nil {
				stack.PushBack(NewCommand("go", command.node.right))
			}

			stack.PushBack(NewCommand("print", command.node))

			if command.node.left != nil {
				stack.PushBack(NewCommand("go", command.node.left))
			}
		}
	}

	return result
}

func PostOrderTraverse(root *TreeNode) []int {
	result := []int{}

	stack := list.New()
	stack.PushBack(NewCommand("go", root))

	for stack.Len() > 0 {
		command := stack.Back().Value.(*Command)
		stack.Remove(stack.Back())

		if command.s == "print" {
			result = append(result, command.node.val)
		} else { // command.s == "go"
			stack.PushBack(NewCommand("print", command.node))

			if command.node.right != nil {
				stack.PushBack(NewCommand("go", command.node.right))
			}

			if command.node.left != nil {
				stack.PushBack(NewCommand("go", command.node.left))
			}
		}
	}

	return result
}
