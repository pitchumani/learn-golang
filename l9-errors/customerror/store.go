package customerror

import (
	"fmt"
	"time"
)

type Model map[string]any

type Store struct {
	data map[string][]Model
}

func (s *Store) All(tn string) ([]Model, error) {
	db := s.data

	if db == nil {
		return nil, fmt.Errorf("no data")
	}

	mods, ok := db[tn]

	if !ok {
		return nil, ErrTableNotFound{
			Table: tn,
			OccurredAt: time.Now(),
		}
	}

	return mods, nil
}

