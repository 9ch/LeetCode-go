package main

import "fmt"

/**
不使用任何内建的哈希表库设计一个哈希集合

具体地说，你的设计应该包含以下的功能

add(value)：向哈希集合中插入一个值。
contains(value) ：返回哈希集合中是否存在这个值。
remove(value)：将给定值从哈希集合中删除。如果哈希集合中没有这个值，什么也不做。

示例:

MyHashSet hashSet = new MyHashSet();
hashSet.add(1);        
hashSet.add(2);        
hashSet.contains(1);    // 返回 true
hashSet.contains(3);    // 返回 false (未找到)
hashSet.add(2);          
hashSet.contains(2);    // 返回 true
hashSet.remove(2);          
hashSet.contains(2);    // 返回  false (已经被删除)

注意：

所有的值都在 [0, 1000000]的范围内。
操作的总数目在[1, 10000]范围内。
不要使用内建的哈希集合库。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/design-hashset
*/

func main() {
	obj := Constructor()
	obj.Add(1)
	obj.Add(2)
	obj.Add(3)
	obj.Remove(1)
	param_3 := obj.Contains(2)
	param_4 := obj.Contains(1)
	param_5 := obj.Contains(3)
	fmt.Println(param_3,param_4,param_5)

}

type MyHashSet struct {
	Mod int
	Arr []interface{}
}

/** Initialize your data structure here. */
func Constructor() MyHashSet {
	hash := MyHashSet{}
	hash.Mod = 1000000
	hash.Arr = make([]interface{}, hash.Mod)
	return hash
}

func (this *MyHashSet) Add(key int) {
	mod := key % this.Mod
	this.Arr[mod] = key
}

func (this *MyHashSet) Remove(key int) {
	mod := key % this.Mod
	if this.Contains(key) {
		this.Arr[mod] = nil
	}
}

/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	mod := key % this.Mod
	if this.Arr[mod] != nil {
		return true
	}
	return false
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 */
