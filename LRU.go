package main

import "fmt"

type Node struct {
	key   int
	value int
	next  *Node
}

// LRUCache data structure
type LRUCache struct {
	capacity int
	cache    map[int]*Node
	head     *Node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{capacity: capacity, cache: make(map[int]*Node)}
}

func (c *LRUCache) Get(key int) int {
	n, exists := c.cache[key]
	if exists {
		c.moveToFront(n)
		return n.value
	}
	return -1
}

func (c *LRUCache) Put(key int, value int) {

	if c.capacity < 1 {
		return
	}

	n, exists := c.cache[key]
	if !exists {
		if len(c.cache) == c.capacity {
			c.removeTail()
		}

		node := Node{key: key, value: value, next: c.head}
		c.cache[key] = &node

		if c.head == nil {
			c.head = &node
			return
		}

		// c.head.prev = &node
		node.next = c.head
		c.head = &node
		return
	}

	n.value = value
	c.moveToFront(n)
}

// -------- Helper methods (you can add these) --------

func (c *LRUCache) moveToFront(node *Node) {
	if c.head.key == node.key {
		return
	}

	currentNode := c.head.next
	previousNode := c.head

	for currentNode != nil && currentNode.key != node.key {
		previousNode = currentNode
		currentNode = currentNode.next
	}

	if currentNode == nil {
		return
	}

	previousNode.next = currentNode.next
	currentNode.next = c.head
	c.head = currentNode
}

func (c *LRUCache) removeTail() {

	if c.head == nil {
		return
	}

	if c.head.next == nil {
		delete(c.cache, c.head.key)
		c.head = nil
		return
	}

	currNode := c.head.next
	previousNode := c.head

	for currNode.next != nil {
		previousNode = currNode
		currNode = currNode.next
	}

	previousNode.next = nil
	delete(c.cache, currNode.key)
}

func main() {

	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Put(1, 10)
	cache.Put(3, 3)

	node := cache.head
	for node != nil {
		fmt.Printf("Node: key %d, value %d \n", node.key, node.value)
		node = node.next
	}

	fmt.Println("--------------------------")

	fmt.Println("get 2", cache.Get(2))

	node = cache.head
	for node != nil {
		fmt.Printf("Node: key %d, value %d \n", node.key, node.value)
		node = node.next
	}

	fmt.Println("--------------------------")

	fmt.Println("get 1", cache.Get(1))

	node = cache.head
	for node != nil {
		fmt.Printf("Node: key %d, value %d \n", node.key, node.value)
		node = node.next
	}

	fmt.Println("--------------------------")

	fmt.Println("get 3", cache.Get(3))

	node = cache.head
	for node != nil {
		fmt.Printf("Node: key %d, value %d \n", node.key, node.value)
		node = node.next
	}
}
