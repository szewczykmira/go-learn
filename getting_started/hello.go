// before go run hello.go run: go mod init hello

package main

import "fmt"

import "rsc.io/quote"

func main() {
    fmt.Println(quote.Opt())
}