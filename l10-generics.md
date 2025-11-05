# Generics

Generic programming is a programming paradigm that allows us to step out
the implementation of a function with a type that will be provided later.
With generics we can write functions that work with multiple types directly
without having to write same function multiple times, one for each type.

## Type Contraints
instead of any type, define constraints

func Name[constraints](parameters) (returns) {
     // ...
}

## definiing constraints

similar to interface, but instead of methods, we need to define types

type MapKey interface {
    int | float64
}

// function has a type constraint K of MapKey, V of any type
//   takes map[K]V as argument
//   returns slice of K type
func Keys[K MapKey, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))

	for k := range m {
		keys = append(keys, k)
	}

	return keys
}

## underlying type constraints

type MyInt int

type MapKey interface {
  int | float64
}

-- this type constraint doesn't allow MyInt though the underlying type is int

to allow types based on the underlying type, use ~ before the type

type MapKey interface {
  ~int | float64
}

see contraint package in golang libraries
e.g.
type Ordered interface {
  Integer | Float | ~string
}

type Integer interface {
  Signed | Unsigned
}

type Signed interface {
  ~int | ~int8 | ~int16 | ~int32 | ~int64
}
...
 
## generic types

