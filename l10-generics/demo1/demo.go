package demo

import "golang.org/x/exp/constraints"

type MyInt int

// define the type constraints
// intersection of int and float64 - allowed types
// used ~ to include underlying type int,
// which allows map[MyInt]string
type MapKey interface {
	~int | float64
}

func Keys[K MapKey, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))

	for k := range m {
		keys = append(keys, k)
	}

	return keys
}


// constraint package has different type constraint defined.
// Ordered is a type constraint for type that are comparable
// it is suitable for the map
func KEYS[K constraints.Ordered, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))

	for k := range m {
		keys = append(keys, k)
	}

	return keys
}
	
