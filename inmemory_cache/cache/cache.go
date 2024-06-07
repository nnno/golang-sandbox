package cache

import "sync"

type Item struct {
	Object interface{}
}

type Cache struct {
	items map[string]Item
	mu    sync.RWMutex
}

func New() *Cache {
	c := &Cache{
		items: make(map[string]Item),
	}
	return c
}

// Set キャッシュに値を追加
func (c *Cache) Set(k string, x interface{}) {
	c.mu.Lock()
	c.items[k] = Item{Object: x}
	c.mu.Unlock()
}

// Get 指定のkeyのデータを取り出す
func (c *Cache) Get(k string) (interface{}, bool) {
	c.mu.RLock()
	item, found := c.items[k]
	if !found {
		c.mu.RUnlock()
		return nil, false
	}
	c.mu.RUnlock()
	return item.Object, true
}

// Delete 指定のkeyのデータを削除
func (c *Cache) Delete(k string) {
	c.mu.Lock()
	delete(c.items, k)
	c.mu.Unlock()
}

// GetAll データを全件取り出す
func (c *Cache) GetAll() map[string]Item {
	c.mu.RLock()
	defer c.mu.RUnlock()
	m := make(map[string]Item)
	for key, val := range c.items {
		m[key] = val
	}
	return m
}
