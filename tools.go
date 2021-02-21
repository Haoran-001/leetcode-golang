package main

import (
	"github.com/cheekybits/genny/generic"
	_ "github.com/cheekybits/genny/generic"
	"sort"
	"sync"
)

func max(x int, y int) int {
	if x >= y {
		return x
	} else {
		return y
	}
}

func min(x int, y int) int {
	if x <= y {
		return x
	} else {
		return y
	}
}

func abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return -x
	}
}

func power(x, y int) int {
	var ans int = 1
	for i := 1; i <= y; i++ {
		ans = ans * x
	}

	return ans
}

func getMaxArrayElements(arr []int) int {
	sort.Ints(arr)
	return arr[len(arr)-1]
}

func getMinArrayElements(arr []int) int {
	sort.Ints(arr)
	return arr[0]
}

func reverseString(s string) string {
	strArr := []byte(s)
	length := len(strArr)

	for i := 0; i < length/2; i++ {
		strArr[i], strArr[length-1-i] = strArr[length-1-i], strArr[i]
	}

	return string(strArr)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Item generic.Type

type ItemStack struct {
	items []Item
	lock  sync.RWMutex
}

// 创建栈
func (s *ItemStack) New() *ItemStack {
	s.items = []Item{}
	return s
}

// 入栈
func (s *ItemStack) Push(t Item) {
	s.lock.Lock()
	s.items = append(s.items, t)
	s.lock.Unlock()
}

// 出栈
func (s *ItemStack) Pop() *Item {
	s.lock.Lock()
	item := s.items[len(s.items)-1]
	s.lock.Unlock()
	return &item
}

type ItemQueue struct {
	items []Item
	lock  sync.RWMutex
}

// 创建队列
func (q *ItemQueue) New() *ItemQueue {
	q.items = []Item{}
	return q
}

// 入队列
func (q *ItemQueue) Enqueue(t Item) {
	q.lock.Lock()
	q.items = append(q.items, t)
	q.lock.Unlock()
}

// 出队列
func (q *ItemQueue) Dequeue() *Item {
	q.lock.Lock()
	item := q.items[0]
	q.items = q.items[1:len(q.items)]
	q.lock.Unlock()
	return &item
}

// 获取队列的第一个元素, 不移除
func (q *ItemQueue) Front() *Item {
	q.lock.Lock()
	item := q.items[0]
	q.lock.Unlock()
	return &item
}

// 判空
func (q *ItemQueue) IsEmpty() bool {
	return len(q.items) == 0
}

// 获取队列的长度
func (q *ItemQueue) Size() int {
	return len(q.items)
}
