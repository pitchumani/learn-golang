package store

import "fmt"
import "golang.org/x/exp/constraints"

type Store[K constraints.Ordered, M Model[K]] struct {
	data map[K]M
}

func (s Store[K, M]) Find(id K) (M, error) {
	m, ok := s.data[id]
	if !ok {
		return m, fmt.Errorf("Key not found %v", id)
	}
	return m, nil
}

func (s Store[K, M]) Insert(m M) error {
	id := m.ID()
	if _, ok := s.data[id]; ok {
		return fmt.Errorf("Key already exists %v", id)
	}
	s.data[id] = m
	return nil
}

