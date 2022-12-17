[![Hits](https://hits.seeyoufarm.com/api/count/incr/badge.svg?url=https%3A%2F%2Fgithub.com%2Fakilans%2Fgolang-mini-projects%2Ftree%2Fmain%2F11-jwt-golang&count_bg=%2379C83D&title_bg=%23555555&icon=&icon_color=%23E7E7E7&title=hits&edge_flat=false)](https://hits.seeyoufarm.com)

# JWT authentication with Golang

- This application uses JWT based authorization method to protect REST api endpoint

### Run - URLs and sample responses

- Run the application with the below commands and access the URLs

  ```bash
  go run main.go
  ```

- Import postman collection "Golang-JWT.postman_collection.json" and start testing

- Access home route - http://localhost:4000 - No auth needed
  ```json
  {
    "status": "Success",
    "message": "Welcome to Golang with JWT authentication"
  }
  ```
- Access secure route - http://localhost:4000/secure - Expect auth error as we didn't pass JWT token

  ```json
  {
    "Status": "Failed",
    "Msg": "You are not authorized to view this page"
  }
  ```

- Generate JWT token - http://localhost:4000/login with payload with the below body

  ```json
  {
    "username": "admin",
    "password": "admin"
  }
  ```

  ```json
  {
    "status": "Success",
    "message": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX25hbWUiOiJBa2lsYW4iLCJMb2dnZWRJblRpbWUiOiIxMC0xMi0yMDIyIDIwOjU3OjMyIiwiaXNzIjoiQWtpbGFuIiwiZXhwIjoxNjcwNjg2MTEyfQ.E--k9nMc-uOHb6VWJCrTyzSgGQ6JGAT_m3J1z_z-Ohs"
  }
  ```

- Copy the JWT token from the above response and pass it in request header as Token value - http://localhost:4000/secure

  ```json
  {
    "status": "Success",
    "message": "Congrats and Welcome to the Secure page!. You gave me the correct JWT token!"
  }
  ```

-

### Demo

![Alt JWT authentication with Golang](https://raw.githubusercontent.com/akilans/golang-mini-projects/main/demos/jwt-with-golang.gif)

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.
