# Maps and Control Structures
## Map
map decl syntax
```go
variable := map[key_type]val_type{}
```
map allows unlimited key value pairs, so cap function can't be used with map

any comparable type can be used as key. Others like func(), map can't be used as keys
Struct type can be used as key when all of its members are of comparable types

iteratable through range (either two values (key, value) or single value (key))
remove elements using delete function

accessing non existing key will return empty value, no error
it is buggy, better to use availability check
```go
value, ok = users["Non existing key"]
if !ok {
}
```
iterating maps in Go is random, can't be in any particular order
retrieve keys to a slice, sort the slice and use the sorted slice
to retrieve the value from map in sorted order

## if statement
if expr {
}

## Switch statement
```go
switch expr {
case val1:
case val2:
default:
}

switch {
case val1 == 1:

case val2 == 2:
}
```
use `fallthrough` at the end of he case body to fall through to next case as well

