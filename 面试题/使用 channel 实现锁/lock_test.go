package main

import (
	"log"
	"testing"
)

func TestMutex(t *testing.T) {
	m := NewMutex()
	sum := 0
	for i := 0; i < 10000; i++ {
		go func() {
			m.Lock()
			sum++
			m.Unlock()
		}()
	}
	if sum == 10000 {
		log.Println("passed")
	} else {
		log.Println("no passed")
	}
}
