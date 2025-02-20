Tutorials and getting started with working with GoLang

- [Instalation](#instalation)
- [Good to know:](#good-to-know)
  - [Commands](#commands)
  - [Syntax](#syntax)
  - [Definitions](#definitions)-
- [First project](getting_started/hello.go)


## Instalation
```sh
brew install go
```


## Good to know:

### Commands

- Initialize module: `go mod init <module_name>` : __module__ allows to track dependencies.
- Run code in file `go run <filename>.go`
-  `go mod tidy` makes sure that mod file matches projects (it's good to use after adding external package)
- `go build` make Go locate the module and add it as a dependency to the go.mod file
- Tests are run with `go test`
- You can also [install app](https://golang.org/doc/tutorial/compile-install)

### Syntax

- `//this is oneline comment`
- Function that name starts with capital letter can be called by a function not in the same package - this is known as `exported name`
- There is `nil` keyword (same as `None`)

- function defintion:
  ```go
  // (type name, type name) return type {
  func FuncName(name variable) int {
    return 42
  }
  ```
- Declaring variables:
  ```go
  var message string
  message = "Hello"
  // or in one line
  information := "Welcome"
  ```

- In Go, code executed as an application must go in a __main__ package.
- you can define custom types:
  ```go
  type fn func(string)(string, error)
  ```

### Definitions

__package__ way to group functions, it's made up of functions in the same directory.