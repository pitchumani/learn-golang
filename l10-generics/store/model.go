package store

import "golang.org/x/exp/constraints"

type Model[T constraints.Ordered] interface {
	ID() T
}

type User struct {
	Email string
}

func (u User) ID() string {
	return u.Email
}
