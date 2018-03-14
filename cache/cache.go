package cache

import (
	"sync"
)

type Cache struct {
	items    map[string]interface{}
	rwMutex  *sync.RWMutex
	capacity int
}

func NewWithCapacity(capacity int) *Cache {
	c := &Cache{}
	c.items = make(map[string]interface{}, capacity)
	c.rwMutex = new(sync.RWMutex)
	c.capacity = capacity
	return c
}

func (p Cache) Add(key string, value interface{}) {
	p.rwMutex.Lock()
	if len(p.items) >= p.capacity {
		for k, _ := range p.items {
			delete(p.items, k)
			break
		}
	}
	p.items[key] = value
	p.rwMutex.Unlock()
}

func (p Cache) Remove(key string) {
	if _, ok := p.Search(key); ok {
		p.rwMutex.Lock()
		delete(p.items, key)
		p.rwMutex.Unlock()
	}
}

func (p Cache) Replace(key string, newValue interface{}) {
	if _, ok := p.Search(key); ok {
		p.rwMutex.Lock()
		p.items[key] = newValue
		p.rwMutex.Unlock()
	}
}

func (p Cache) Search(key string) (interface{}, bool) {
	p.rwMutex.RLock()
	v, found := p.items[key]
	p.rwMutex.RUnlock()
	return v, found
}
