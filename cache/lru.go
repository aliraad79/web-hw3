package main

type lru struct {
}

func (l *lru) evict(c *Cache) string {
	key := c.doublyLinkedList.Front().key
	c.doublyLinkedList.RemoveFromFront()
	return key
}

func (l *lru) get(node *node, c *Cache) {
	c.doublyLinkedList.MoveNodeToEnd(node)
}

func (l *lru) set(node *node, c *Cache) {
	c.doublyLinkedList.AddToEnd(node)
}

func (l *lru) set_overwrite(node *node, value string, c *Cache) {
	node.value = value
	c.doublyLinkedList.MoveNodeToEnd(node)
}
