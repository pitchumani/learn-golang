package demo

import (
	"sync"
)
/*
type Store struct {
	data map[int]bool
	mu   sync.Mutex
}

func (s *Store) Add(i int) {
	s.mu.Lock()
	s.data[i] = true
	s.mu.Unlock()
}

func (s *Store) Get(i int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.data[i]
}

*/
type Store struct {
	data map[int]bool
	mu   sync.RWMutex
}

func (s *Store) Add(i int) {
	s.mu.Lock()
	s.data[i] = true
	s.mu.Unlock()
}

func (s *Store) Get(i int) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.data[i]
}
