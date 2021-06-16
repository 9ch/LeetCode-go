package main

import (
	"fmt"
	"log"
)

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func main() {

}

func Pre(root *TreeNode) []int {
	var result []int

	var dfs func(r *TreeNode)
	dfs = func(r *TreeNode) {
		if r != nil {
			result = append(result, r.Val)
			dfs(r.Left)
			dfs(r.Right)
		}
	}
	dfs(root)
	return result
}

/**
使用迭代栈来实现前序遍历
*/
func Pre2(root *TreeNode) []int {
	s := &Stack{}
	var result []int
	if root == nil {
		return result
	}
	s.Push(root)
	for !s.IsEmpty() {
		n, _ := s.Pop()
		result = append(result, n.Val)
		if n.Right != nil {
			s.Push(n.Right)
		}
		if n.Left != nil {
			s.Push(n.Left)
		}
	}

	return result
}

/**
堆
*/
type Stack struct {
	data   []interface{}
	length int64
}

func (s *Stack) Push(v interface{}) bool {
	if s.IsEmpty() {
		log.Println("堆为空")
		return false
	}
	s.data[s.length] = v
	s.length++
	return true
}

func (s *Stack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		log.Println("堆为空")
		err := fmt.Errorf("堆为空")
		return nil, err
	}
	s.length--
	return s.data[s.length], nil

}

/**
返回堆是否为空
*/
func (s *Stack) IsEmpty() bool {
	return s.length == 0
}
