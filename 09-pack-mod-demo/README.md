[![Hits](https://hits.seeyoufarm.com/api/count/incr/badge.svg?url=https%3A%2F%2Fgithub.com%2Fakilans%2Fgolang-mini-projects%2Ftree%2Fmain%2F09-pack-mod-demo&count_bg=%2379C83D&title_bg=%23555555&icon=&icon_color=%23E7E7E7&title=hits&edge_flat=false)](https://hits.seeyoufarm.com)

# Golang Package and Module demo

This program uses package from [https://github.com/akilans/go-module-demo]

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
