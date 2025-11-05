# Arrays, Slices and Iterations

- two builtin collection types arrays and slices
- allows cmoposition to create more complex data structures, e.g. list
- array is fixed in size, slice is not
- slices can wrap arrays
- garbage collected after use

- no member functions like size, append, remove
- generics can be used to implement such function by end users (append, len, cap)
- cap is different from len, it is dependent on go implementation in specific arch, os


## iterations
- for loop is similar to C (pre, condition and post op)
- len(collection) can be used in for loop
- range can be used for collection types array and slice. it returns index and value for each iteration. can use only index with range (e.g. for i := range names)


