package main

type Cache struct {
	list map[int]int
}

func LRUcache(capacity int) *Cache {
	return &Cache{
		list: make(map[int]int, capacity),
	}
}

func (c *Cache) Get() int {

	return -1
}

func (c *Cache) Put(key, value int) {
	if _, ok := c.list[key]; ok {
		c.list[key] = value
	}

	c.list[key] = value
}
