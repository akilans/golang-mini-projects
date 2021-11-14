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
# VS code installs default tools needed for golang syntax
# But good to know how it works
# fmt goimports
go fmt . # automatically reformats your code to match the standard format.
# Install 3rd party packages
# check in /home/akilan/go-workspace/bin & pkg dir
go install golang.org/x/tools/cmd/goimports@latest

goimports -l -w .
#The -l flag tells goimports to print the files with incorrect formatting to the console.
#The -w flag tells goimports to modify the files in-place.
#The . specifies the files to be scanned: everything in the
#current directory and all of its subdirectories.
# golint and go vet
# Refer https://golangci-lint.run/ for Golang CI
```

Rather than use separate tools, you can run
multiple tools together with golangci-lint. It combines golint, go
vet, and an ever-increasing set of other code quality tools. Once it is
installed, you run golangci-lint with the command

By default vs uses staticcheck

## Useful commands

```bash
# remove path from terminal display
export PS1="\[\e]0;\u@\h: \w\a\]${debian_chroot:+($debian_chroot)}\[\033[01;32m\]\u@\h\[\033[00m\]:\$ "
```
