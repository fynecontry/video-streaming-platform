package main

import (
	"fmt"
)

type LFUCache struct {
	capacity       int
	size           int
	minFrequency   int
	cacheFrequency map[int]*LinkedList
	cache          map[string]*LinkedListNode
}

func (lf *LFUCache) Get(key string) *LinkedListNode {
	if _, ok := lf.cache[key]; !ok {
		return nil
	}
	temp := lf.cache[key]

	lf.cacheFrequency[temp.frequency].RemoveNode(temp)
	if lf.cacheFrequency[lf.cache[key].frequency].head == nil {
		delete(lf.cacheFrequency, lf.cache[key].frequency)
		if lf.minFrequency == lf.cache[key].frequency {
			lf.cache[key].frequency++
		}
	}
	lf.cache[key].frequency++

	if _, ok := lf.cacheFrequency[lf.cache[key].frequency]; !ok {
		lf.cacheFrequency[lf.cache[key].frequency] = &LinkedList{}
	}
	lf.cacheFrequency[lf.cache[key].frequency].Append(lf.cache[key])
	return lf.cache[key]
}

func (lf *LFUCache) Set(key string, value int) {
	if lf.Get(key) != nil {
		lf.cache[key].rank = value
		return
	}
	if lf.size == lf.capacity {
		delete(lf.cache, lf.cacheFrequency[lf.minFrequency].head.show)
		lf.cacheFrequency[lf.minFrequency].RemoveNode(lf.cacheFrequency[lf.minFrequency].head)
		if lf.cacheFrequency[lf.minFrequency].head == nil {
			delete(lf.cacheFrequency, lf.minFrequency)
		}
		lf.size--
	}
	lf.minFrequency = 1
	lf.cache[key] = &LinkedListNode{show: key, rank: value, frequency: lf.minFrequency}
	if _, ok := lf.cacheFrequency[lf.cache[key].frequency]; !ok {
		lf.cacheFrequency[lf.cache[key].frequency] = &LinkedList{}
	}
	lf.cacheFrequency[lf.cache[key].frequency].Append(lf.cache[key])
	lf.size++
}

func (lf *LFUCache) Print() {
	for first, second := range lf.cache {
		fmt.Printf("(%v,%v)", first, second.rank)
	}
	fmt.Printf("\n\n")
}
