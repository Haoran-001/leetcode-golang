package main

import (
	"sync"
)

type IntStack struct {
	items []int
	lock  sync.RWMutex
}

// 创建栈
func (s *IntStack) New() *IntStack {
	s.items = []int{}
	return s
}

// 入栈
func (s *IntStack) Push(t int) {
	s.lock.Lock()
	s.items = append(s.items, t)
	s.lock.Unlock()
}

// 出栈
func (s *IntStack) Pop() *int {
	s.lock.Lock()
	item := s.items[len(s.items)-1]
	s.items = s.items[0 : len(s.items)-1]
	s.lock.Unlock()
	return &item
}
