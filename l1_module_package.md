## Module
A Go module is a collection of packages and their dependencies that can be built, versioned, and versioned as a unit.

Go toolchain has support for underlying VCS systems such as Git and Mercurial. Go shall work
with the modules hosted on these systems.
e.g. to host a module in github initialize module like:
go mod init github.com/user/repo

to use a module hosted in github
e.g.
go get github.com/gobuffalo/flect

this will add the flect module as dependencies in your current module
go.sum file will be generated, it will have all direct and indirect dependencies
by the our module.

go get -u package -- will update the package dependecies. using without package name will
update all the dependent packages in the module.
go mod tidy -- cleanup the old dependencies etc

semantic versioning: MAJOR.MINOR.PATCH
major version change should update the import path as well.
e.g.
"github.com/gofrs/uuid"
"github.com/gofrs/uuid/v5"

avoid cyclical import

use alias name to avoid name collison
e.g. both bar and foo/bar used in a file
     "demo/bar"
     "demo/foo/bar"

use like below:
    "demo/bar"
    pub "demo/foo/bar"


## Package
A package is a collection of Go files that share the same import path. These
files are always in the same directory as each other, and the directory name is
almost always the same as the package name.

Packages can export variables, functions or types. Exported items shall be referenced
by other packages. Unexported items are accessible within the package only.

naming rules:
- lower-cased
- only letters and numbers
- starts with letter
- must be declared at the top of every go source file

e.g. package http, package sql
avoid using builtin package names

File, Folder organization:
less strict rules.
good names: lower case and underscore (e.g. service_manager.go, user.go, user_test.go)
bad names: ServiceManagerTest.go, User.go, UserTest.go

make the package for single purpose, e.g. http and time separately

Dependencies




