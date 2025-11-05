# Testing

## basics

tests are in Go itself.
file named <source>_test.go is identified as test file

test function
- must be in _test.go file
- must have prefix "Test"
- must accept a *testing.T type
- must return no arguments

testing.T type has few methods to mark test as failure, cleanup resources,
run subset of tests etc

.Fatal()
.Error()

## running tests

go test

- run all tests recursively
go test ./...

- run tests in specific package
go test ./mdx

- use -v option to print verbose

- use .Log method of testing.T type to log details from the test

- package tests run in separate thread, the individual tests within packages are not
- use t.Parallel() method to run tests in parallel

- use -run "regex" option to run matching tests
- use -timeout 50ms option to limit the execution time, avoid infinite loops

## code coverage
- use -cover option with go test to report code coverage
- coverprofile=coverage.out - report
go

### table driven testing

### test helpers
- setup, teardown or provision resources for the tests
- can be used to write assertions or mock out external dependencies
- it is a normal Go function with testing.TB as first argument
- call .Helper() from helper function to notify Go that it is helper function
- use t.Cleanup() to do tasks like removing the file generated etc, a test helper
  can have multiple Cleannup calls


