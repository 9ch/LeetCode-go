package main

import (
	"fmt"
	"sync"
)

////支持add
//
////支持delete
//
////支持迭代

type ISortedMap interface {
	Len() int                            //map 的长度
	Add(key, value string)               //添加
	Delete(key string)                   //删除
	Next() func() (string, string, bool) //迭代
}
type SortedMap struct {
	sort []string          //存储有序的结构
	data map[string]*Entry //存储当前索引的位置以及值
	mu   sync.RWMutex
}

// Entry 。。。记录在数组里面的索引位置，方便后面 O（1）操作获取 key 删除
type Entry struct {
	Key   int
	Value string
}

func NewMap() *SortedMap {
	return &SortedMap{data: map[string]*Entry{}}
}

// Len 获取当前map 的长度。
func (s *SortedMap) Len() int {
	return len(s.sort)
}

// Add 添加操作，
func (s *SortedMap) Add(key, value string) {
	s.mu.Lock()
	s.mu.Unlock()
	if _, ok := s.data[key]; !ok {
		s.data[key] = &Entry{
			Key:   s.Len(),
			Value: value,
		}
		s.sort = append(s.sort, key)
	} else {
		s.data[key].Value = value
	}

}

func (s *SortedMap) Delete(key string) {
	if v, ok := s.data[key]; ok {
		s.sort = append(s.sort[:v.Key], s.sort[v.Key+1:]...)
	}
	delete(s.data, key)
}

func (s *SortedMap) Next() func() (string, string, bool) {
	i := 0
	return func() (key string, value string, bool2 bool) {
		s.mu.RLock()
		defer s.mu.RUnlock()
		if i >= s.Len() {
			return
		}
		key, value = s.sort[i], s.data[s.sort[i]].Value
		i++
		return key, value, true
	}
}

func main() {
	m := NewMap()

	m.Add("k1", "v1")
	m.Add("k2", "v2")
	m.Add("k3", "v3")
	f := m.Next()
	for {
		k, v, ok := f()
		if !ok {
			break
		}
		fmt.Println(k, v)
	}
	m.Delete("k3")
	f = m.Next()
	for {
		k, v, ok := f()
		if !ok {
			break
		}
		fmt.Println(k, v)
	}
}
