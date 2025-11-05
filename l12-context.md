# Context

Context package was introduced in go 1.7, provides the cleaner way than the
use of channels for managing cancellations and timeouts across the go routines.

Context is mostly used for controlling concurrent systems in the applications.

- contexts
- context values
- cancellation propagation
- context errors

## Contexts

Context is an interface, has four methods (Deadline, Done, Err, Value)

Deadline - check if a context has a cancellation deadline set
Done - listen for cancellation events - similar to listening to channel being closed
Err - check if context is cancelled
Value - get the value from the context


not recommended to store the context
pass the context as function arguments, usually as first arg of the function

## cancellation propagation

the cancellation of a context, will cancel all of its children context as well

## context err

- context canceled error
- context deadline exceeded error
- cancel context with custom error

