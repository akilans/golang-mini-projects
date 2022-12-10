package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// Store the SECRET KEY SECRETLY :)
var SECRET_KEY string = "AwesomeGolangSecret"

// To capture credentials from request which is needed to generate JWT
type User struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

// message to send as json response
type Message struct {
	Status string `json:"status"`
	Msg    string `json:"message"`
}

// response message struct as json format
func jsonMessageByte(status string, msg string) []byte {
	errrMessage := Message{status, msg}
	byteContent, _ := json.Marshal(errrMessage)
	return byteContent
}

// Custom claims needed for generating JWT token
type MyCustomClaims struct {
	UserName     string `json:"user_name"`
	LoggedInTime string
	jwt.RegisteredClaims
}

// Function to create JWT token
func CreateJWT() (string, error) {
	currentTime := time.Now().Format("02-01-2006 15:04:05")

	// Storing user name and loggedin time
	// Token expires in 1 min. This is just for testing
	claims := MyCustomClaims{
		"Akilan",
		currentTime,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Minute)),
			Issuer:    "Akilan",
		},
	}

	// Generate token with HS256 algorithm and custom claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Sign the token with our secret key
	signedToken, err := token.SignedString([]byte(SECRET_KEY))

	return signedToken, err
}

// Function to validate JWT
// Get the token from user and validate it
func ValidateJWT(tokenString string) bool {
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRET_KEY), nil
	})

	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		log.Printf("%v - %v - %v \n", claims.UserName, claims.LoggedInTime, claims.RegisteredClaims.Issuer)
		return true
	} else {
		log.Println(err)
		return false
	}
}

// Middleware auth handler
func Auth(handler func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the JWT token from request header
		if r.Header["Token"] != nil {
			providedToken := r.Header["Token"][0]
			if ValidateJWT(providedToken) {
				handler(w, r)
			} else {
				w.WriteHeader(401)
				w.Write(jsonMessageByte("Failed", "You are not authorized to view this page"))
			}
		} else {
			w.WriteHeader(401)
			w.Write(jsonMessageByte("Failed", "Please provide valid JWT token in request header as Token"))
		}

	})
}

// Handle login
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(405)
		w.Write(jsonMessageByte("Failed", r.Method+" - Method not allowed"))
	} else {
		var userData User
		err := json.NewDecoder(r.Body).Decode(&userData)
		if err != nil {
			w.WriteHeader(400)
			w.Write(jsonMessageByte("Failed", "Bad Request - Failed to parse the payload "))
		} else {
			log.Printf("User name - %v and Password is %v\n", userData.UserName, userData.Password)
			// user name and password is hard code
			// We can use DB
			if userData.UserName == "admin" && userData.Password == "admin" {
				token, _ := CreateJWT()
				w.Write(jsonMessageByte("Success", token))
			} else {
				w.WriteHeader(401)
				w.Write(jsonMessageByte("Failed", "Invalid credentials"))
			}
		}
	}

}

// Handle home route
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write(jsonMessageByte("Success", "Welcome to Golang with JWT authentication"))
}

// Handle secure route
func SecureHandler(w http.ResponseWriter, r *http.Request) {
	w.Write(jsonMessageByte("Success", "Congrats and Welcome to the Secure page!. You gave me the correct JWT token!"))
}

func main() {
	fmt.Println("JWT - authentication with Golang")

	// No auth needed
	http.HandleFunc("/", HomeHandler)

	// Generate JWT token by providing username and password in the payload
	http.HandleFunc("/login", LoginHandler)

	// Auth middleware added for restricing the direct acccess
	// Provide JWT token in header section as "Token"
	http.HandleFunc("/secure", Auth(SecureHandler))

	http.ListenAndServe(":4000", nil)
}
