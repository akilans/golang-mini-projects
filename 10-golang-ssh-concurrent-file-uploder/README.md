# Concurrent File Uploader to remote servers using Go Routine

It is golang command line application to upload files to remote servers concurrently using golang's inbuilt go routine and waitgroup

## Demo

![Alt Concurrent SSH File Uploader](https://raw.githubusercontent.com/akilans/golang-mini-projects/main/demos/golang-ssh-concurrent-file-uploder.gif)

## Usage

- Normal method will upload hello.txt file to all remote servers
- Go routine will upload hi.txt file to all remote servers
- hello.txt and hi.txt file size is same

```bash

# clone a repo
git clone https://github.com/akilans/golang-mini-projects.git

# go to the 10-golang-ssh-concurrent-file-uploder dir
cd 10-golang-ssh-concurrent-file-uploder

# build
go build

# run

./golang-ssh-concurrent-file-uploder

# sample oputput
./golang-ssh-concurrent-file-uploder
File hello.txt uploaded successfully to 127.0.0.1:2222
File hello.txt uploaded successfully to 127.0.0.1:2200
File hello.txt uploaded successfully to 127.0.0.1:2201
File hello.txt uploaded successfully to 127.0.0.1:2202
Without Go Routine - Upload task took 4.13 seconds
File hi.txt uploaded successfully to 127.0.0.1:2222
File hi.txt uploaded successfully to 127.0.0.1:2202
File hi.txt uploaded successfully to 127.0.0.1:2201
File hi.txt uploaded successfully to 127.0.0.1:2200
With Go Routine - Upload task took 2.19 seconds
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.
