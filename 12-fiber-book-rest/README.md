[![Hits](https://hits.seeyoufarm.com/api/count/incr/badge.svg?url=https%3A%2F%2Fgithub.com%2Fakilans%2Fgolang-mini-projects%2Ftree%2Fmain%2F12-fiber-book-rest&count_bg=%2379C83D&title_bg=%23555555&icon=&icon_color=%23E7E7E7&title=hits&edge_flat=false)](https://hits.seeyoufarm.com)

# Bookstore REST API with MySQL, GORM, JWT and Fiber framework

- This is REST based API to list, add, update and delete books

### Tools and Packages

- Fiber - Golang web Framework
- MySql - SQL Database
- Gorm - ORM library for golang
- JWT - For Authorization
- Bycrypt - To hash passwords

### Prerequisites

- Golang
- Mysql

```bash
# I used docker to run mysql
docker container run --name mysql -d -p 3306:3306 -e MYSQL_ROOT_PASSWORD=root#123 mysql
docker exec -it mysql bash
mysql -u root -p
Enter password: root#123
mysql> CREATE DATABASE bookstore;
```

- Update .env file with correct information

```bash
DB_DSN="root:root#123@tcp(127.0.0.1:3306)/bookstore?charset=utf8mb4&parseTime=True&loc=Local"
PORT=":8080"
SECRET_KEY="mysecret"
```

### Run the application

```bash
go run main.go
```

### URL endpoints

- Users
  - `POST` - Register - /admin
  - `POST` - Login - /login
- Books - JWT token needs to be passed in header
  - `GET` - Get all the books - `/`
  - `GET` - Get book by id - `/book/{id}`
  - `DELETE` - Delete book by id - `/book/{id}`
  - `POST` - Add a book - `/addbook`
  - `PUT` - Update book - `/book/{id}`

### Postman setup for testing

- Import `fiber-rest-book.postman_collection.json` and start testing

### Demo

![Alt Bookstore API](https://raw.githubusercontent.com/akilans/golang-mini-projects/main/demos/fiber-book-rest.gif)
