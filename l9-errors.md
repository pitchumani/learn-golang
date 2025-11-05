# Errors

## errors
- errors are implemented as error interface

type error interface {
     Error() string
}

## panic
panic - fatal error
recover - builtin function
defer and recover can be used to handle the panics gracefully

capturing and returning panic values:
- let panic crash app and deal with fall out
- recover from panic, log it, and move on
- capture panicked value and return it as an error

Fixing the panic
- use a defer with a recover to catch the panic
- type assertion on the value returned from the panic to see if it was an error
- used a named return to allow sending back of the error from the deferred recover


You shouldn't raise panic from a package because user should have control
of the program, package shouldn't dictate the control flow.
Panic may be appropriate from main as user has control over it.

Few panic situations:
uninitialized map, type assertion, out of bounds slice array access

runtime debug package provides useful functions for debugging
use Stack() function to get stack trace when required
PrintStack() prints stack trace into std out

## custom errors

## wrapping/ unwrapping errors
errors package provides Unwrap function to unwrap the errors

## using errors.As and errors.Is
