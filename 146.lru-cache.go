package main

import "fmt"

func main() {
	cache := Constructor(2)
}

type Node struct {
	Key        int
	Val        int
	prev, next *Node
}

type DList struct {
	head, tail *Node
}

func (dl *DList) appendLast(n *Node) {
	pr := dl.tail.prev
	pr.next = n
	n.prev = pr
	n.next = dl.tail
	dl.tail.prev = n
}

func (dl *DList) remove(n *Node) {
	n.prev.next = n.next
	n.next.prev = n.prev
}

func (dl *DList) tranverse() {
	p := dl.head.next
	for p != dl.tail {
		fmt.Print(p.Val, "->")
		p = p.next
	}
}

type LRUCache struct {
	table    map[int]*Node
	dl       *DList
	capacity int
}

func Constructor(capacity int) LRUCache {
	n1, n2 := &Node{}, &Node{}
	n1.next = n2
	n2.prev = n1
	return LRUCache{make(map[int]*Node), &DList{head: n1, tail: n2}, capacity}
}

func (this *LRUCache) Get(key int) int {
	if n, ok := this.table[key]; ok {
		this.dl.remove(n)
		this.dl.appendLast(n)
		return n.Val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if n, ok := this.table[key]; ok {
		this.dl.remove(n)
		this.dl.appendLast(n)
		n.Val = value
	} else {
		if len(this.table) == this.capacity {
			n := this.dl.head.next
			if n == this.dl.tail {
				return
			}
			this.dl.remove(n)
			delete(this.table, n.Key)
		}

		newNode := &Node{Key: key, Val: value}
		this.dl.appendLast(newNode)
		this.table[key] = newNode
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
