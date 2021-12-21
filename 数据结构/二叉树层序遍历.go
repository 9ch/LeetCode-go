package main

import (
	"container/list"
	"fmt"
)

//二叉树的层序遍历

func main() {
	t := TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}
	fmt.Println(levelLoop(&t))
}

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func levelLoop(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}
	queue := list.New()
	queue.PushFront(root)
	for queue.Len() > 0 {
		length := queue.Len()
		var current []int
		for i := 0; i < length; i++ {
			ele := queue.Remove(queue.Back()).(*TreeNode)
			current = append(current, ele.Val)
			if ele.Left != nil {
				queue.PushFront(ele.Left)
			}
			if ele.Right != nil {
				queue.PushFront(ele.Right)
			}
		}
		result = append(result, current)
	}
	return result
}
