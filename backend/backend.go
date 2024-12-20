package backend

import "sync"

type Monster struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Size string `json:"size"`
}

type Backend struct {
	mu       sync.Mutex
	monsters map[int]Monster
}

// NewBackend creates a new Backend
func NewBackend() *Backend {
	return &Backend{
		monsters: make(map[int]Monster),
	}
}

// Set stores a value in the Backend
func (b *Backend) Set(key int, value Monster) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.monsters[key] = value
}

// Get retrieves a value from the Backend
func (b *Backend) Get(key int) (Monster, bool) {
	b.mu.Lock()
	defer b.mu.Unlock()
	val, ok := b.monsters[key]
	return val, ok
}
