# Go language basics
## overview
compiled language
aimed to compile fast

statically typed, garbage collecting compiled language
capable of producing concurrent thread-safe programs

executable are statically linked, self contained.
i.e. that don't need any runtime and dev libraries to run

## Numbers
signed/ unsigned integers

int8, uint8
int16, uint16
int32, uint32
int64, uint64

float32
float64

complex64
complex128

byte - alias for uint8
rune - alias for int32

implementation specific types
uint - either 32 or 64 bits (32bit vs 64bit machines)
int - same size uint
uintptr -

## Strings and UTF-8
backquotes - raw strings, contains any character except backtick, multi line strings
doublequotes - interpreted strings - with escape seq

var a rune = 'A'
fmt.Printf("%q %v (%T)\n", a, a, a)

use range instead len to iterate string characters as multibyte chars will be treated differently.

utf-8 can be used in program itself, as identifier or function names.

## Variable
statically typed, known at compile time

var <type> varname
varname = init_val

var <type> varname = init_val

varname := init_val  // type is inferred from init value

varname := type(init_val) // to enforce specific type


no public, private for visibility
instead capitalization is used in Golang
if an identifier starts with uppercase letter, that identifier is accessible outside the package, it is called exported

printing - check fmt package, close to C lang print
* %d - integer
* %5d - minimal places to the number
* %05d - prefixed with zeros if num is less than 5 digits
* %T - print type of the variable
* %v - value of the variable
* %+v - more info to the value of the variable (e.g. struct var with all field names)
* %#v - go syntax format e.g. main.User{Name: "Kurt", Age:32}
* %[i]s - specifier for i'th argument - specifier and argument can be in diff order

strconv package can be used to convert string.
strconv.ParseInt("-42", 0, 64)
strconv.ParseFloat("42.123456", 64)
strconv.Atoi("42")


