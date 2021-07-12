package main

import (
	"fmt"
	"hash/crc32"
)

//LRU内存淘汰算法
//通过使用 hash 和双端链表的方式来实现，hash 用来快速取值，链表用来快速定位

type Lru struct {
	Data   map[uint32]*ListNode //定义数据存放
	Length int
}

//定义一个双端链表
type ListNode struct {
	Val       string
	Next, Pre *ListNode
}

var (
	Tail *ListNode
	Head *ListNode
)

func main() {
	lru := NewLru(3)
	lru.Put("hello", "123")
	lru.Put("world", "321")
	lru.Put("aw", "fawa")
	fmt.Println(lru)

}

func NewLru(length int) Lru {
	return Lru{
		Data:   make(map[uint32]*ListNode),
		Length: length,
	}
}

// Put 添加元素
func (lru *Lru) Put(key, data string) error {

	hash := crc32.ChecksumIEEE([]byte(key))
	if _, ok := lru.Data[hash]; ok {
		if lru.Data[hash] != Tail {
			lru.refresh(lru.Data[hash])
		}
	} else {
		list := &ListNode{
			Val:  data,
			Next: nil,
			Pre:  Tail,
		}
		lru.Data[hash] = list
		Tail = list
		if lru.Length > 0 {
			lru.Length--
		} else {
			lru.Remove()
		}
		lru.refresh(list)
	}
	return nil
}

// Get 获取元素
func (lru *Lru) Get(key string) (string, error) {

	hash := crc32.ChecksumIEEE([]byte(key))
	if v, ok := lru.Data[hash]; !ok {
		return "", fmt.Errorf("%s", "该元素不存在！")
	} else {
		result := v.Val
		lru.refresh(v)
		return result, nil
	}
}

// 重置最新访问的节点。
func (lru *Lru) refresh(node *ListNode) {
	if node == Tail {
		return
	}
	if Head == nil {
		Head = node
	} else {
		node.Pre.Next = node.Next
		node.Next.Pre = node.Pre
		Tail.Next = node
		node.Pre = Tail
		Tail = node
	}

}

//将未使用的元素删除
func (lru *Lru) Remove() {
	hash := crc32.ChecksumIEEE([]byte(Head.Val))
	delete(lru.Data, hash)
	Head = Head.Next
	Head.Pre = nil
	lru.Length++
}
