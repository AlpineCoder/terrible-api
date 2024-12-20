package store

import (
	"sync"

	"github.com/AlpineCoder/terrible-api/business/models"
)

type Store struct {
	Mu        sync.Mutex
	Monsters  map[int]models.Monster
	IdCounter int
}

// NewBackend creates a new Backend
func NewBackend() *Store {
	return &Store{
		Monsters: make(map[int]models.Monster),
	}
}

// Set stores a value in the Backend
func (b *Store) Set(key int, value models.Monster) {
	b.Mu.Lock()
	defer b.Mu.Unlock()
	b.Monsters[key] = value
}

// Get retrieves a value from the Backend
func (b *Store) Get(key int) (models.Monster, bool) {
	b.Mu.Lock()
	defer b.Mu.Unlock()
	val, ok := b.Monsters[key]
	return val, ok
}

func (b *Store) NextID() int {
	b.Mu.Lock()
	defer b.Mu.Unlock()
	b.IdCounter = b.IdCounter + 1
	return b.IdCounter
}
