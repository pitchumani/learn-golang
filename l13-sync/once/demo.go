package demo

import (
	"fmt"
	"sync"
	"time"
)

/*
type Builder struct {
	Built bool
}

func (b *Builder) Build() error {
	fmt.Println("building...")

	time.Sleep(10 * time.Millisecond)

	fmt.Println("built")

	b.Built = true

	return nil
}
*/

type Builder struct {
	Built bool
	Once  sync.Once
}

func (b *Builder) Build() error {
	b.Once.Do(func() {
		fmt.Println("building...")

		time.Sleep(10 * time.Millisecond)

		fmt.Println("built")

		b.Built = true
	})
	return nil
}
