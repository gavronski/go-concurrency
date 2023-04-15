package deadLock

import "sync"

var m sync.Mutex

type Collection struct {
	Data map[string]string
}

func NewCollection() Collection {
	return Collection{
		Data: make(map[string]string),
	}
}
func (c *Collection) Add(key, value string) {
	m.Lock()
	defer m.Unlock()

	if c.Has(key) {
		return
	}

	c.Data[key] = value
}

func (c *Collection) Has(key string) bool {
	m.Lock()
	defer m.Unlock()

	if _, ok := c.Data[key]; ok {
		return true
	}

	return false
}
