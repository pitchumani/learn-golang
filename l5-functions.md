# Functions

func name(argname type)(ret_type,...) {
}

all return values must be captured, unwanted
return values can be ignored using _ placeholder

named return: return without return value can return the value of the named return.

func f1()(retval bool) {
     retval = false
     return
}

named return is useful to differentiate same type return values
to be careful, see example of defer function

first class functions
functions in Go are types, can be passed as argument for another function

anonymous function - function without name

variadic arguments - must be last argument. (arg1 int, names ...sting)

deferring function calls - defer keyword - defer the execution of a function until the return of the parent function
- useful for handling cleanup tasks
- deferred panic is executed at the end
- however, the arguments to defer functions are evaluated immediately

init function
- intended to initialize the package
- takes no arguments, returns nothing
- executed when the containing package is imported or executed
- multiple init functions can be created; they are executed in the order of their definitions
- if multiple packages has init, they are executed in the order of their imports
