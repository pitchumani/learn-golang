# structs, methods and pointers

## structs
- custom complex type declaration
- contains zero or more fields, 0 or more methods
- type keyword is used for struct type declaration
- multi line initialization require field name to be specified
- single line doesn't require field name, order matters, all field initial values must

- struct tag (a metadata) can be used with each field after type definition
- declared inside backtick
e.g. Name string `json:"name,omitempty"`
- other packages can make use of struct tag info
- package encoding/json uses struct tag 'json'


## methods
it is syntactic sugar
when declaring a function, before the function name, the set of parenthesis are
introduced to allow the user to define the receiver of the method.
it is similar to functions, but can't be called independently, it should be
called with the receiver object.

function type can be declared using type keyword

## pointers

A type that holds address of another variable.
In Go, Variables are passed by value into other functions.
There is no pass-by-reference in Go. Pointers are used to simulate that similar to C.

User & or new function to create pointer
