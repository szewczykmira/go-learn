# go-learn
Tutorials and getting started with working with GoLang


## Good to know:

### Syntax

- `//this is oneline comment`

- Initialize module: `go mod init <module_name>`
- Run code in file `go run <filename>.go`
-  `go mod tidy` makes sure that mod file matches projects

- Function that name starts with capital letter can be called by a function not in the same package - this is known as `exported name`
- `go build` make Go locate the module and add it as a dependency to the go.mod file
- There is `nil` keyword (same as `None`)