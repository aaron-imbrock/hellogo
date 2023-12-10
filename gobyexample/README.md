# Go by Example

## hello-world.go

```go
package main

import "fmt"

func main() {
    fmt.Println("hello world")
}
```

## Run is IDE, build is Compile

```bash
$ go run hello-world.go
hello world
```

Sometimes we’ll want to build our programs into binaries. We can do this using `go build`.

```bash
$ go build hello-world.go
$ ls
hello-world    hello-world.go
```

We can then execute the built binary directly.

```bash
$ ./hello-world
hello world
```