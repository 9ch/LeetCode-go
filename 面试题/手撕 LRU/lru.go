package main

import (
	"container/list"
)

//LRU内存淘汰算法
//通过使用 hash 和双端链表的方式来实现，hash 用来快速取值，链表用来快速定位

type CallBack func(key string, value Value)

type iLru interface {
	Get(key string) (Value, bool) //Get 方法
	Add(key string, value Value)  // Add 方法
}

type Lru struct {
	Data     map[string]*list.Element //定义数据存放
	MaxBytes int64                    //最大字节数
	NowBytes int64                    //当前字节数
	ll       *list.List               //双端链表
	Callback CallBack                 //当元素被移除时的回调函数
}

type Value interface {
	Len() int //返回该数据的长度
}

// 链表当中书籍存放的数据
type entry struct {
	key   string // 做一份冗余，可以快读的定位 map
	value Value  //实际存放的值，每个值都需要实现 Len 方法
}

func NewLru(max int64, callback CallBack) *Lru {
	return &Lru{
		Data:     make(map[string]*list.Element),
		MaxBytes: max,
		NowBytes: 0,
		ll:       list.New(),
		Callback: callback,
	}
}

// 获取元素
func (l *Lru) Get(key string) (value Value, ok bool) {
	if ele, ok := l.Data[key]; ok {
		l.ll.MoveToFront(ele)
		return ele.Value.(*entry).value, true
	}
	return
}

//添加元素
func (l *Lru) Add(key string, value Value) {
	if ele, ok := l.Data[key]; ok {
		l.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		l.NowBytes += int64(value.Len()) - int64(kv.value.Len())
		ele.Value = &entry{key, value}
	} else {
		ele := l.ll.PushFront(&entry{key, value})
		l.Data[key] = ele
		l.NowBytes += int64(len(key)) + int64(value.Len())
	}
	if l.MaxBytes != 0 && l.NowBytes > l.MaxBytes {
		l.remove()
	}
}

//删除元素
func (l *Lru) remove() {
	ele := l.ll.Back()
	if ele != nil {
		l.ll.Remove(ele)
		kv := ele.Value.(*entry)
		delete(l.Data, kv.key)
		l.NowBytes -= int64(len(kv.key)) + int64(kv.value.Len())
		if l.Callback != nil {
			l.Callback(kv.key, kv.value)
		}
	}
}

// 获取 lru 里面元素的数目
func (l *Lru) Len() int {
	return l.ll.Len()
}

var _ iLru = (*Lru)(nil)
