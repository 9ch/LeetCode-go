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
	fmt.Println(Pre2(&t))
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
	s := NewS()
	var result []int
	if root == nil {
		return result
	}
	s.Push(root)
	fmt.Println(s.data)
	for !s.IsEmpty() {
		t, _ := s.Pop()
		n := t.(*TreeNode)
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

func NewS() *Stack {
	return &Stack{
		data:   make([]interface{}, 0),
		length: 0,
	}
}
func (s *Stack) Push(v interface{}) bool {
	s.data = append(s.data, v)
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
	t := s.data[s.length]
	s.data = s.data[:s.length]
	return t, nil

}

/**
返回堆是否为空
*/
func (s *Stack) IsEmpty() bool {
	return s.length == 0
}
