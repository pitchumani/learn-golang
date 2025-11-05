# synchronization

- wait groups
- error groups
- data races
- locking with a mutex
- performing tasks only once

## Wait groups
## error groups
## data races
When two or more goroutines try to access the same shared resource.

go test -race -- will warn about data races
will be slow as it does lot of checking to find the data races

## locking with mutex

The sync locker interface defines methods that a type must implement
in order to be able to lock and unlock a shared resource.
methods

sync mutex type is Mutex available in go

sync.Mutex is heavy weight implementation, only one go routine is allowed
to access the shared resource
mu sync.Mutex

mu.Lock()
mu.Unlock()

sync.RWMutex is lighter version, that can be used to allow multiple go routines
to read from shared resource, but allow only one goroutine to write.

mu sync.RWMutex

mu.Lock()
mu.Unlock()
mu.RLock()
mu.RUnlock()

## performing tasks only once
Once sync.Once

Once.Do(func() {
  ...
})

useful for tasks such as closing channels
(closing the same channel more than once will cause panic)
