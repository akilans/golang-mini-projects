package main

import (
	"fmt"
	// bye is a folder name and goodbye is a package name.
	// Best practice is use the same name for folder and package like hello package
	goodbye "github.com/akilans/go-module-demo/bye"
	"github.com/akilans/go-module-demo/hello"
)

func main() {
	fmt.Println("Hello")
	hello.SayHello() // SayHello exported by default as it follows camelcase pattern
	goodbye.SayBye() // SayBye exported by default as it follows camelcase pattern
}
