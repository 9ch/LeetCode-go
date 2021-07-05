package main

import "fmt"

//给定一个只包括 '('，')'，'{'，'}'，'['，']'的字符串 s ，判断字符串是否有效。
//
//有效字符串需满足：
//
//左括号必须用相同类型的右括号闭合。
//左括号必须以正确的顺序闭合。
//
//
//示例 1：
//
//输入：s = "()"
//输出：true
//示例2：
//
//输入：s = "()[]{}"
//输出：true
//示例3：
//
//输入：s = "(]"
//输出：false
//示例4：
//
//输入：s = "([)]"
//输出：false
//示例5：
//
//输入：s = "{[]}"
//输出：true
//
//
//提示：
//
//1 <= s.length <= 104
//s 仅由括号 '()[]{}' 组成
//
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/valid-parentheses

func main() {
	s := "()"
	fmt.Println(isValid(s))
	s2 := "()[]{}"
	fmt.Println(isValid(s2))
	s3 := "(]"
	fmt.Println(isValid(s3))
	s4 := "{[)]"
	fmt.Println(isValid(s4))
	s5 := "{[]}"
	fmt.Println(isValid(s5))
}

func isValid(s string) bool {

	length := len(s)

	//如果长度不是偶数的话，则肯定不符合要求，直接返回 false
	if length%2 != 0 {
		return false
	}
	//定义一个符号的对应表
	m := make(map[rune]rune, 3)
	m[')'] = '('
	m['}'] = '{'
	m[']'] = '['
	stack := Stack{data: make([]interface{}, length)}
	//每次判断栈内的元素是否跟当前元素相对应，如果不对应，则直接返回 false
	for _, v := range s {
		if v == '(' || v == '{' || v == '[' {
			_ = stack.push(v)
		} else if va := stack.put(); va != m[v] {
			return false
		}
	}
	if stack.length > 0 {
		return false
	}
	return true
}

type Stack struct {
	data   []interface{}
	length int
}

func (s *Stack) put() interface{} {
	result := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	s.length--
	return result
}

func (s *Stack) push(value interface{}) error {
	s.data = append(s.data, value)
	s.length++
	return nil
}
