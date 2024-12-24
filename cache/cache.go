package cache

import "fmt"

type Cache struct {
	items map[string]interface{}
}

func NewCache() *Cache {
	return &Cache{
		items: make(map[string]interface{}),
	}
}

func (c *Cache) Set(key string, value interface{}) {
	prev_value := c.items[key]
	if prev_value != nil {
		c.Delete(key)
		fmt.Println("Удаленно")
	}
	c.items[key] = value
	fmt.Println("Добавленно")
}

func (c *Cache) Get(key string) interface{} {
	return c.items[key]
}

func (c *Cache) Delete(key string) {
	delete(c.items, key)
}
