# Golang Package and Module demo

### Notes

- Have you ever used any packages in simple go app?. Definitely yes
- Below example import "fmt" package. Do you wonder why all the functions, variable follows Camelcase pattern?
- It is because fmt package exports variables, functions to other packages.
- If any variable or function not following Camelcase then it will not be exported. It will be used inside a parent package

```go
package main

import (
    "fmt"
)

func main(){
    fmt.Println("Hello Go!")
    fmt.Printf("Hello Go!\n")
}

```

### Reference Module

- In this application we are using the below module
- [Please refer the go-module-demo here ](https://github.com/akilans/go-module-demo)

### Run this app

```bash
go mod tidy
go run main.go
```
