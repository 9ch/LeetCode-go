package main

import (
	"fmt"
	"time"
)

//使用 channel 实现锁

type Mutex struct {
	ch chan struct{}
}

func NewMutex() *Mutex {
	mu := &Mutex{ch: make(chan struct{}, 1)}
	mu.ch <- struct{}{}
	return mu
}
func (m *Mutex) Lock() {
	<-m.ch
}

func (m *Mutex) Unlock() bool {
	select {
	case m.ch <- struct{}{}:
		return true
	default:
		return false
	}
}

func (m *Mutex) TryLock() bool {
	select {
	case <-m.ch:
		return true
	default:
	}
	return false
}

func main() {
	m := NewMutex()
	sum := 0
	for i := 0; i < 10000; i++ {
		go func() {
			m.Lock()
			sum++
			m.Unlock()
		}()
	}
	time.Sleep(time.Second)
	fmt.Println(sum)
}

// output
// 10000
