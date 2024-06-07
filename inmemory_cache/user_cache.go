// インメモリでのキャッシュ機構
// sync.Mutexを使って排他制御を行う
package main

import "sync"

// User ユーザー情報
type User struct {
	ID   int
	Name string
}

type (
	UserSlice []User
)

// UserCache キャッシュ
type UserCache struct {
	data map[string]User
	mu   sync.Mutex
}

// NewUserCache Cacheを生成
func NewUserCache() *UserCache {
	c := &UserCache{data: make(map[string]User)}
	return c
}

// Get 指定のkeyのデータを取り出す
func (c *UserCache) Get(key string) (User, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.data[key], true
}

// GetAll データを全件取り出す
func (c *UserCache) GetAll() UserSlice {
	c.mu.Lock()
	defer c.mu.Unlock()

	var ret []User
	for _, v := range c.data {
		ret = append(ret, v)
	}
	return ret
}

// Set データをキャッシュする
func (c *UserCache) Set(key string, x User) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, ok := c.data[key]; !ok {
		c.data[key] = x
	}
}
