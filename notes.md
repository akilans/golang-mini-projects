# Golang Notes for Beginners

### Install Golang on Linux

```bash
# Download go1.17.3.linux-amd64.tar.gz
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.17.3.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
# Create workspace folder anywhere and update here
`
go-workspace/
├── bin
├── pkg
└── src
    └── github.com
        └── akilans
            └── go-tutorial
                ├── 01-hello
                │   └── main.go
                └── readme.md
`
export GOPATH=/home/akilan/go-workspace
export PATH=$PATH:$GOPATH/bin
go version
```

### Go Formatting

```bash
# instead of default name of binary give the custom one
go build -o hello_world hello.go

# go has no central repo like maven, npm. It depends on SCM
# VS code installs default tools needed for golang syntax
# But good to know how it works
# fmt goimports
go fmt . # automatically reformats your code to match the standard format.
# Install 3rd party packages
# check in /home/akilan/go-workspace/bin & pkg dir
go install golang.org/x/tools/cmd/goimports@latest

# cleans up your import statements. It puts them in
#alphabetical order, removes unused imports, and attempts to guess
#any unspecified imports
goimports -l -w .
#The -l flag tells goimports to print the files with incorrect formatting to the console.
#The -w flag tells goimports to modify the files in-place.
#The . specifies the files to be scanned: everything in the
#current directory and all of its subdirectories.
# golint and go vet
# Refer https://golangci-lint.run/ for Golang CI
```

go lint - checks for any syntax error
go vet - check for any unused variable, invalid function param etc.

Always run go fmt or goimports before compiling your code!

Rather than use separate tools, you can run
multiple tools together with golangci-lint. It combines golint, go
vet, and an ever-increasing set of other code quality tools. Once it is
installed, you run golangci-lint with the command

By default vs code uses staticcheck, Delve, gopls
Delve is a debugger for the Go programming language.

## Useful commands

```bash
# remove path from terminal display
export PS1="\[\e]0;\u@\h: \w\a\]${debian_chroot:+($debian_chroot)}\[\033[01;32m\]\u@\h\[\033[00m\]:\$ "
ffmpeg -i ~/test.mp4 ~/test.gif
```

## Notes

- If you are referring to a character, use the rune type
- When you are within a function, you can use the := operator to replace a var declaration

```go
//different ways to declare and assign variable
var x int = 10
var x = 10
var x int
var x, y int = 10, 20
var x, y int
var x, y = 10, "hello"
var (
    x int
    y = 20
    z int = 30
    d, e = 40, "hello"
    f, g string
)
```

When you take a slice from a slice, you are not making a copy of the
data. Instead, you now have two variables that are sharing memory.
This means that changes to an element in a slice affect all slices that
share that element.

```go
x := [] int {1, 2, 3, 4}
y := x[:2]
z := x[1:]x[1] = 20
y[0] = 10
z[1] = 30
fmt.Println("x:", x)
fmt.Println("y:", y)
fmt.Println("z:", z)

x: [10 20 30 4]
y: [10 20]
z: [20 30 4]

```

Be very careful when taking a slice of a slice! Both slices share the
same memory and changes to one are reflected in the other. Avoid
modifying slices after they have been sliced or if they were produced by
slicing. Use a three-part slice expression to prevent append from sharing
capacity between slices.

```go
// copy whole slice
x := [] int {1, 2, 3, 4}
y := make([] int , 4)
num := copy(y, x)
fmt.Println(y, num)

// [1 2 3 4] 4

// copy few elements from slice
x := [] int {1, 2, 3, 4}
y := make([] int , 2)
copy(y, x[2:])

// another example
x := [] int {1, 2, 3, 4}
d := [4] int {5, 6, 7, 8}
y := make([] int , 2)
copy(y, d[:])
fmt.Println(y)
copy(d[:], x)
fmt.Println(d)
// output
[5 6]
[1 2 3 4]
```
