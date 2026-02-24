package lru

// doublyLinkedNode represents a node in the doubly linked list used for LRU ordering.
// The most recently used item is at the front, least recently used at the back.
type doublyLinkedNode struct {
	key      string
	value    int
	previous *doublyLinkedNode
	next     *doublyLinkedNode
}

// Cache implements a Least Recently Used cache with O(1) get and put operations.
// Uses a hash map for fast lookups and a doubly linked list for LRU ordering.
type Cache struct {
	capacity    int
	lookup      map[string]*doublyLinkedNode
	headSentinel *doublyLinkedNode
	tailSentinel *doublyLinkedNode
}

// NewCache creates an LRU cache with the given maximum capacity.
func NewCache(capacity int) *Cache {
	// Sentinel nodes simplify edge cases for list operations
	headSentinel := &doublyLinkedNode{}
	tailSentinel := &doublyLinkedNode{}
	headSentinel.next = tailSentinel
	tailSentinel.previous = headSentinel

	return &Cache{
		capacity:    capacity,
		lookup:      make(map[string]*doublyLinkedNode),
		headSentinel: headSentinel,
		tailSentinel: tailSentinel,
	}
}

// Get retrieves a value from the cache and marks it as recently used.
// Returns the value and true if found, 0 and false otherwise.
func (c *Cache) Get(key string) (int, bool) {
	node, exists := c.lookup[key]
	if !exists {
		return 0, false
	}

	// Move the accessed node to the front (most recently used position)
	c.removeNode(node)
	c.addToFront(node)

	return node.value, true
}

// Put adds or updates a key-value pair in the cache.
// If the cache is full, the least recently used item is evicted.
func (c *Cache) Put(key string, value int) {
	// If the key already exists, update value and move to front
	if existingNode, exists := c.lookup[key]; exists {
		existingNode.value = value
		c.removeNode(existingNode)
		c.addToFront(existingNode)
		return
	}

	// Evict the least recently used item if cache is at capacity
	if len(c.lookup) >= c.capacity {
		leastRecentNode := c.tailSentinel.previous
		c.removeNode(leastRecentNode)
		delete(c.lookup, leastRecentNode.key)
	}

	// Create and add the new node
	newNode := &doublyLinkedNode{key: key, value: value}
	c.addToFront(newNode)
	c.lookup[key] = newNode
}

// Size returns the current number of items in the cache.
func (c *Cache) Size() int {
	return len(c.lookup)
}

// removeNode detaches a node from the doubly linked list.
func (c *Cache) removeNode(node *doublyLinkedNode) {
	node.previous.next = node.next
	node.next.previous = node.previous
}

// addToFront places a node right after the head sentinel (most recently used position).
func (c *Cache) addToFront(node *doublyLinkedNode) {
	node.next = c.headSentinel.next
	node.previous = c.headSentinel
	c.headSentinel.next.previous = node
	c.headSentinel.next = node
}
