package demo

type Store struct {
	data map[int]bool
}

func (s *Store) Add(i int) {
	s.data[i] = true
}

func (s *Store) Get(i int) bool {
	return s.data[i]
}

