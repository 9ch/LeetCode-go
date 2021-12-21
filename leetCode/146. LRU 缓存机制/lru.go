package main

import (
	"container/list"
	"fmt"
)

//运用你所掌握的数据结构，设计和实现一个  LRU (最近最少使用) 缓存机制 。
//实现 LRUCache 类：
//
//LRUCache(int capacity) 以正整数作为容量 capacity 初始化 LRU 缓存
//int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
//void put(int key, int value) 如果关键字已经存在，则变更其数据值；如果关键字不存在，则插入该组「关键字-值」。当缓存容量达到上限时，它应该在写入新数据之前删除最久未使用的数据值，从而为新的数据值留出空间。
//
//
//进阶：你是否可以在 O(1) 时间复杂度内完成这两种操作？
//
//
//
//示例：
//
//输入
//["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
//[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
//输出
//[null, null, null, 1, null, -1, null, -1, 3, 4]
//
//解释
//LRUCache lRUCache = new LRUCache(2);
//lRUCache.put(1, 1); // 缓存是 {1=1}
//lRUCache.put(2, 2); // 缓存是 {1=1, 2=2}
//lRUCache.get(1);    // 返回 1
//lRUCache.put(3, 3); // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
//lRUCache.get(2);    // 返回 -1 (未找到)
//lRUCache.put(4, 4); // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
//lRUCache.get(1);    // 返回 -1 (未找到)
//lRUCache.get(3);    // 返回 3
//lRUCache.get(4);    // 返回 4
//
//
//提示：
//
//1 <= capacity <= 3000
//0 <= key <= 10000
//0 <= value <= 105
//最多调用 2 * 105 次 get 和 put
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/lru-cache

type LRUCache struct {
	data   map[int]*list.Element
	ll     *list.List
	MaxCap int //最大容量
	NowCap int //当前容量

}

// 缓存里面实际存放的值，冗余一份 key，方便后期查找
type entry struct {
	Key   int
	Value int
}

//返回Lru
func Constructor(capacity int) LRUCache {
	return LRUCache{
		data:   make(map[int]*list.Element),
		ll:     list.New(),
		MaxCap: capacity,
		NowCap: 0,
	}
}

// Get get 操作
func (this *LRUCache) Get(key int) int {
	if v, ok := this.data[key]; ok {
		this.ll.MoveToFront(v)
		ele := v.Value.(*entry)
		e := ele.Value
		return e
	}
	return -1
}

// Put  set 操作
func (this *LRUCache) Put(key int, value int) {
	if v, ok := this.data[key]; ok {
		this.ll.MoveToFront(v)
		v.Value = &entry{
			Key:   key,
			Value: value,
		}
	} else {
		tmp := &entry{
			Key:   key,
			Value: value,
		}
		ele := this.ll.PushFront(tmp)
		this.data[key] = ele
		this.NowCap++
	}
	if this.MaxCap != 0 && this.NowCap > this.MaxCap {
		this.remove(key)
	}
}

// 删除列表的最后一个元素
func (this *LRUCache) remove(key int) {
	ele := this.ll.Back()
	if ele != nil {
		e := ele.Value.(*entry)
		delete(this.data, e.Key)
		this.ll.Remove(ele)
		this.NowCap--
	}
}
func (l *LRUCache) Len() int {
	return len(l.data)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
func main() {
	lru := Constructor(3)
	lru.Put(1, 1)
	lru.Put(2, 2)
	lru.Put(3, 3)
	fmt.Println(lru.Get(1))
	lru.Put(4, 4)
	fmt.Println(lru.Get(4))
	fmt.Println(lru.Get(2))
	fmt.Println(lru.Len())
}
