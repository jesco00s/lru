package main

import "fmt"

type Node struct {
	key   int
	value int
	prev  *Node
	next  *Node
}

// LRUCache data structure
type LRUCache struct {
	capacity int
	cache    map[int]*Node
	head     *Node // Most recently used
	tail     *Node // Least recently used
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

		c.head.prev = &node
		node.next = c.head
		c.head = &node
		return
	}

	n.value = value
	c.moveToFront(n)
}

// -------- Helper methods (you can add these) --------

// removeNode removes a node from the doubly linked list.
func (c *LRUCache) removeNode(node *Node) {
	if c.head.key == node.key {
		delete(c.cache, node.key)
		if c.head.next == nil {
			c.head = nil
		} else {
			c.head = c.head.next
		}
	}

	//TODO need to add logic to remove node if not the head
	//iterate through until we find the node
}

// addToFront inserts a node right after head.
func (c *LRUCache) addToFront(node *Node) {
	// TODO
}

func (c *LRUCache) moveToFront(node *Node) {
	if c.head.key == node.key {
		return
	}

	currentNode := c.head.next
	previousNode := c.head

	for currentNode.key != node.key {
		previousNode = currentNode
		currentNode = currentNode.next
	}

	previousNode.next = nil
	c.head.prev = currentNode
	currentNode.next = c.head
	currentNode.prev = nil
	c.head = currentNode
}

func (c *LRUCache) removeTail() {
	tailNode := c.head.next
	previousNode := c.head

	for tailNode.next != nil {
		previousNode = tailNode
		tailNode = tailNode.next
	}

	previousNode.next = nil
	delete(c.cache, tailNode.key)
}

func main() {

	cache := Constructor(2)

	cache.Put(1, 100)

	cache.Put(2, 200)

	node := cache.head
	for node != nil {
		fmt.Printf("Node: key %d, value %d \n", node.key, node.value)
		node = node.next
	}

	fmt.Println("--------------------------")

	cache.Put(1, 101)

	node = cache.head
	for node != nil {
		fmt.Printf("Node: key %d, value %d \n", node.key, node.value)
		node = node.next
	}

	fmt.Println("--------------------------")

	cache.Get(2)

	node = cache.head
	for node != nil {
		fmt.Printf("Node: key %d, value %d \n", node.key, node.value)
		node = node.next
	}

	fmt.Println("--------------------------")

	cache.Put(3, 300)

	node = cache.head
	for node != nil {
		fmt.Printf("Node: key %d, value %d \n", node.key, node.value)
		node = node.next
	}

	fmt.Println("--------------------------")

	cache.Put(4, 400)

	node = cache.head
	for node != nil {
		fmt.Printf("Node: key %d, value %d \n", node.key, node.value)
		node = node.next
	}

	// cache.removeNode(&Node{key: 1, value: 100})
}
