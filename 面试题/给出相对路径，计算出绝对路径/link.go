package main

import (
	"fmt"
	"strings"
)

// a/b/c/../../d  给出相对路径 实现绝对路径

func main() {
	fmt.Println(link("a/b/c/../../d"))
}

//通过使用栈的方式
func link(s string) string {
	if len(s) <= 1 {
		return s
	}
	var stack1 []string
	var result string
	t := strings.Split(s, "/")
	fmt.Println(t)
	for i := 0; i < len(t); i++ {
		if t[i] != ".." {
			stack1 = append(stack1, t[i])
		} else {
			stack1 = stack1[:len(stack1)-1]
		}
	}
	for i := 0; i < len(stack1); i++ {
		result += stack1[i] + "/"
	}
	return result
}
