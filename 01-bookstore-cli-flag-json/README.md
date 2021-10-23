# Book Store Cli

It is golang command line application to list, add, update and delete books using flag, json, ioutils package

## Demo

![Alt Organize Folder](https://raw.githubusercontent.com/akilans/golang-mini-projects/main/demos/golang-bookstore-cli.gif)

## Usage

```bash

# clone a repo
git clone https://github.com/akilans/golang-mini-projects.git

# go to the 01-bookstore-cli-flag-json dir
cd 01-bookstore-cli-flag-json

# build
go build

# run

# get books
./bookstore get --all
./bookstore get --id 5

# add a book with id ,title, author, price, image_url
./bookstore add --id 6 --title test-book --author akilan --price 200 --image_url http://akilan.com/test.png

# update a book with id ,title, author, price, image_url
./bookstore update --id 6 --title test-book-1 --author akilan1 --price 2001 --image_url http://akilan.com/test.png1

# delete a book by --id
./bookstore delete --id 6

```

## Credits and references

1. [That DevOps Guy](https://www.youtube.com/c/MarcelDempers)
2. [Donald Feury](https://www.youtube.com/c/DonaldFeury)

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.
