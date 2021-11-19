[![Hits](https://hits.seeyoufarm.com/api/count/incr/badge.svg?url=https%3A%2F%2Fgithub.com%2Fakilans%2Fgolang-mini-projects%2Ftree%2Fmain%2F05-random-password-flag&count_bg=%2379C83D&title_bg=%23555555&icon=&icon_color=%23E7E7E7&title=hits&edge_flat=false)](https://hits.seeyoufarm.com)

# Generate Random Passwords using Golang

It is golang command line application to generate random passwords
using "math/rand" and flag golang inbuilt package

## Demo

![Alt Generate random password](https://raw.githubusercontent.com/akilans/golang-mini-projects/main/demos/golang-random-password-flag.gif)

## Usage

```bash

# clone a repo
git clone https://github.com/akilans/golang-mini-projects.git

# go to the 05-random-password dir
cd 05-random-password-flag

# build
go build

# run

./random-password

# Enter the number of passwords you want to generate
# sample oputput
./random-password --count=3 --length=10 --min-number=2 --min-special=2 --min-upper=2
Password 1 is bC#09z3v55
Password 2 is f2311wm5M-
Password 3 is 9$q660h6hH

```

## Reference

[Golang By Example](https://golangbyexample.com/generate-random-password-golang/)

## Credits

[Jackson Atkins](https://github.com/realugbun)

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.
