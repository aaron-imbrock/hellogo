# Learning Go Reading Notes

Created Dec 09, 2023 by Aaron Imbrock.

## Table of Contents

- [Setting Up Your Go Environment](ch01/README.md)
- [Primitive Types and Declarations](ch02/README.md)
- [Composite Types](ch03/README.md)
- [Blocks, Shadows, and Control Structures](ch04/README.md)
- [Functions](ch05/README.md)
- [Pointers](ch06/README.md)
- [Types, Methods, and Interfaces](ch07/README.md)
- [Errors](ch08/README.md)
- [Modules, Packages, and Imports](ch09/README.md)
- [Concurrency in Go](ch10/README.md)
- [The Standard Library](ch11/README.md)
- [The Context](ch12/README.md)
- [Writing Tests](ch13/README.md)
- [Here There Be Dragons](ch14/README.md)
- [Into the Future: Generics in Go](ch15/README.md)

## Go Command

- `go run ch01/hello.go` - Temporarily builds and then executes `hello.go` and runs it.
- `go build ch01/hello.go` - Compiles executable called `hello`.
- `go build -o hello_world ch01/hello.go` - Compiles `hello.go` into executable `hello_world`.

## Go Third-Party Tools

`go install github.com/rakyll/hey@latest`

```bash
GOPATH defaults to $HOME/go/bin
Installs `hey` to $GOPATH/bin
```

