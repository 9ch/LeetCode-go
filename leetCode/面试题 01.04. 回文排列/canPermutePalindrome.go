package main

//给定一个字符串，编写一个函数判定其是否为某个回文串的排列之一。
//
//回文串是指正反两个方向都一样的单词或短语。排列是指字母的重新排列。
//
//回文串不一定是字典当中的单词。
//
//
//
//示例1：
//
//输入："tactcoa"
//输出：true（排列有"tacocat"、"atcocta"，等等）

func main() {

}

//回文排列
//
func canPermutePalindrome(s string) bool {

	//定义一个临时存放数量的 map
	tempMap := make(map[string]int)

	//将每个字符让如 map 中，并计数
	for _, v := range s {
		tempMap[string(v)]++
	}
	//如果出现两个字符以上的次数为奇数则说明不成立
	i := 0
	for _, v := range tempMap {
		if v%2 != 0 {
			if i%2 != 0 {
				return false
			}
			i++
		}
	}

	return true
}
