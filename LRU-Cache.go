package main

import (
	"fmt"
)

type LRUCache struct {
	capacity int

	//LinkedListNode holds key and value pairs
	cache     map[string]*LinkedListNode
	cacheVals LinkedList
}

func (lr *LRUCache) Get(key string) *LinkedListNode {
	if _, ok := lr.cache[key]; !ok {
		return nil
	} else {
		value := lr.cache[key].rank
		lr.cacheVals.RemoveNode(lr.cache[key])
		lr.cacheVals.InsertAtTail(key, value)
		return lr.cacheVals.GetTail()
	}
}

func (lr *LRUCache) Set(key string, value int) {
	if _, ok := lr.cache[key]; !ok {
		if lr.cacheVals.size >= lr.capacity {
			lr.cacheVals.InsertAtTail(key, value)
			lr.cache[key] = lr.cacheVals.GetTail()
			delete(lr.cache, lr.cacheVals.GetHead().show)
			lr.cacheVals.RemoveHead()
		} else {
			lr.cacheVals.InsertAtTail(key, value)
			lr.cache[key] = lr.cacheVals.GetTail()
		}
	} else {
		lr.cacheVals.RemoveNode(lr.cache[key])
		lr.cacheVals.InsertAtTail(key, value)
		lr.cache[key] = lr.cacheVals.GetTail()
	}
}

func (lr *LRUCache) Print() {
	curr := lr.cacheVals.head
	for curr != nil {
		fmt.Printf("(%v,%v)", curr.show, curr.rank)
		curr = curr.next
	}
	fmt.Printf("\n\n")
}
