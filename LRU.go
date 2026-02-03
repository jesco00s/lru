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
	return LRUCache{capacity: capacity}
}

// Get retrieves a value from the cache.
// Moves the accessed node to the head (most recent).
func (c *LRUCache) Get(key int) int {
	// TODO
	n, exists := c.cache[key]
	if exists {
		return n.value
	}
	return -1
}

// Put adds or updates a value.
// If at capacity, remove the tail (least recent).
func (c *LRUCache) Put(key int, value int) {
	// TODO figure out if you need to remove a node before inserting

	//Need to check if key exists in the cache
	//If exists remove old one and insert a new one at the head.

	//else make a new node
	//insert at the head
	node := Node{key: key, value: value, next: c.head}
	c.head = &node
	c.cache[key] = &node

}

// -------- Helper methods (you can add these) --------

// removeNode removes a node from the doubly linked list.
func (c *LRUCache) removeNode(node *Node) {
	// TODO
}

// addToFront inserts a node right after head.
func (c *LRUCache) addToFront(node *Node) {
	// TODO
}

// moveToFront moves an existing node to the front.
func (c *LRUCache) moveToFront(node *Node) {
	// TODO
}

// removeTail removes the tail node and returns it.
// func (c *LRUCache) removeTail() *Node {
// 	// TODO
// }

func main() {

	cache := Constructor(2)

	cache.Put(1, 100)
	fmt.Println("cache", cache)
}
