[![Hits](https://hits.seeyoufarm.com/api/count/incr/badge.svg?url=https%3A%2F%2Fgithub.com%2Fakilans%2Fgolang-mini-projects%2Ftree%2Fmain%2F07-ssh-sftp-agent&count_bg=%2379C83D&title_bg=%23555555&icon=&icon_color=%23E7E7E7&title=hits&edge_flat=false)](https://hits.seeyoufarm.com)

# Execute command, Create, Upload and Download file using Golang

It is a golang based application to execute, create, upload, and download a file to and from a remote server using ssh and sftp package

This can improved by flag cli but my aim is to make you understand SSH and SFTP logic with golang

## Prerequisites

- Go
- SSH server with username and password authentication or key based authentication

## Demo

![Alt SSH and SFTP agent](https://raw.githubusercontent.com/akilans/golang-mini-projects/main/demos/golang-ssh-sftp.gif)

## Usage

This program will do the following tasks

- Execute a command on remote server
- Create a file on remote server
- Upload a local file to remote sever
- Download a remote file to local

```bash

# clone a repo
git clone https://github.com/akilans/golang-mini-projects.git

# go to the 07-ssh-agent dir
cd 07-ssh-sftp-agent

# build
go build

# run

./go-ssh-sftp


```
