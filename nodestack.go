package main

import "sync"

type NodeStack struct {
	NodeItems []*Node
	lock      sync.RWMutex
}

func (stack *NodeStack) Push(item *Node) {
	stack.lock.Lock()
	stack.NodeItems = append(stack.NodeItems, item)
	stack.lock.Unlock()
}

func (stack *NodeStack) Pop() *Node {
	stack.lock.Lock()
	item := stack.NodeItems[stack.Len()-1]
	stack.NodeItems = stack.NodeItems[:stack.Len()-1]
	stack.lock.Unlock()

	return item
}

func (stack *NodeStack) Len() int {
	return len(stack.NodeItems)
}

func (stack *NodeStack) Empty() bool {
	return stack.Len() == 0
}
