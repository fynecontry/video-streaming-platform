package main

import (
	"fmt"
)

type LinkedListNode struct {
	show      string
	rank      int
	frequency int
	next      *LinkedListNode
	prev      *LinkedListNode
}

type LinkedList struct {
	head *LinkedListNode
	tail *LinkedListNode
	size int
}

func createLinkedList(data [][]interface{}) *LinkedListNode {
	var head *LinkedListNode
	var tail *LinkedListNode
	for _, pair := range data {
		key, value := pair[0].(string), pair[1].(int)
		newNode := &LinkedListNode{show: key, rank: value}
		if head == nil {
			head = newNode
		} else {
			tail.next = newNode
		}
		tail = newNode
	}
	return head
}

func (l *LinkedList) InsertAtHead(key string, value int) {
	newNode := &LinkedListNode{show: key, rank: value}
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		newNode.next = l.head
		l.head.prev = newNode
		l.head = newNode
	}
	l.size++
}

func (l *LinkedList) InsertAtTail(key string, value int) {
	newNode := &LinkedListNode{show: key, rank: value}
	if l.tail == nil {
		l.tail = newNode
		l.head = newNode
		newNode.next = nil
	} else {
		l.tail.next = newNode
		newNode.prev = l.tail
		l.tail = newNode
		newNode.next = nil
	}
	l.size++
}

func (l *LinkedList) GetHead() *LinkedListNode {
	return l.head
}

func (l *LinkedList) GetTail() *LinkedListNode {
	return l.tail
}

func (l *LinkedList) RemoveNode(node *LinkedListNode) *LinkedListNode {
	if node == nil {
		return nil
	}

	if node.prev != nil {
		node.prev.next = node.next
	}

	if node.next != nil {
		node.next.prev = node.prev
	}

	if node == l.head {

		l.head = l.head.next
	}
	if node == l.tail {
		l.tail = l.tail.prev
	}
	l.size--
	return node
}

func (l *LinkedList) Remove(key string) {
	i := l.GetHead()
	for i != nil {
		if i.show == key {
			l.RemoveNode(i)
		}
		i = i.next
	}
}

func (l *LinkedList) RemoveHead() *LinkedListNode {
	return l.RemoveNode(l.head)
}

func (l *LinkedList) RemoveTail() *LinkedListNode {
	return l.RemoveNode(l.tail)
}

func (l *LinkedList) Append(node *LinkedListNode) {
	if l.head == nil {
		l.head = node
	} else {
		l.tail.next = node
		node.prev = l.tail
	}
	l.tail = node
}

func display(head *LinkedListNode) {
	temp := head
	for temp != nil {
		fmt.Printf("%s\t\tRatings: %d\n", temp.show, temp.rank)
		temp = temp.next
		if temp != nil {
			continue
		}
	}
	fmt.Printf("\n")
}

func merge2SortedLists(l1, l2 *LinkedListNode) *LinkedListNode {
	dummy := &LinkedListNode{rank: -1}
	prev := dummy
	for l1 != nil && l2 != nil {
		if l1.rank <= l2.rank {
			prev.next = l1
			l1 = l1.next
		} else {
			prev.next = l2
			l2 = l2.next
		}
		prev = prev.next
	}

	if l1 == nil {
		prev.next = l2
	} else {
		prev.next = l1
	}

	return dummy.next
}

func MergeKSortedLists(lists []*LinkedListNode) *LinkedListNode {
	if len(lists) > 0 {
		res := lists[0]

		for i := 1; i < len(lists); i++ {
			res = merge2SortedLists(res, lists[i])
		}
		return res
	}

	return &LinkedListNode{rank: -1}
}
