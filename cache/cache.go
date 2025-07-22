package cache

import "sync"

// Interface - реализуйте этот интерфейс
type Interface interface {
	Set(k, v string)
	Get(k string) (v string, ok bool)
}

// Не меняйте названия структуры и название метода создания экземпляра Cache, иначе не будут проходить тесты

type Cache struct {
	data map[string]string
	mu   sync.RWMutex
}

// NewCache создаёт и возвращает новый экземпляр Cache.
func NewCache() Interface {
	m := make(map[string]string)
	return &Cache{data: m}
}

func (c *Cache) Set(k, v string) {
	c.mu.Lock()
	c.data[k] = v
	c.mu.Unlock()
}

func (c *Cache) Get(k string) (v string, ok bool) {
	c.mu.RLock()
	v, ok = c.data[k]
	c.mu.RUnlock()
	
	return
}
