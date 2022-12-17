package controllers

import (
	"os"
	"strconv"
	"time"

	"github.com/akilans/fiber-book-rest/helpers"
	"github.com/akilans/fiber-book-rest/models"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type Login struct {
	Email    string `json:"email" validate:"required,email,min=6,max=100"`
	Password string `json:"password" validate:"required,min=6,max=15"`
}

// Custom claims needed for generating JWT token
type MyCustomClaims struct {
	UserEmail    string
	LoggedInTime string
	jwt.RegisteredClaims
}

// validate user payload
func ValidateUserStruct(user models.User) []ErrorResponse {
	var errors []ErrorResponse
	err := validate.Struct(user)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, element)
		}
	}
	return errors
}

func ValidateLoginUserStruct(user Login) []ErrorResponse {
	var errors []ErrorResponse
	err := validate.Struct(user)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, element)
		}
	}
	return errors
}

// Add User Handler
/*
Get email and password from user
Validate email and password based on Login struct rules
hash the password using bcrypt
Store email and hashedpassword in db
*/
func AddUserHandler(c *fiber.Ctx) error {
	var loginUser models.User

	//parse the request body and store the email and password inputs
	if err := c.BodyParser(&loginUser); err != nil {
		helpers.LogError(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Please provide valid inputs",
		})
	} else {
		// validate user inputs
		errors := ValidateUserStruct(loginUser)
		if errors != nil {
			return c.Status(fiber.StatusBadRequest).JSON(errors)
		}

		var existingUser models.User

		existingUser, err = models.GetUserByEmail(loginUser.Email)

		if err != nil {
			helpers.LogError(err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed get user info",
			})
		}

		if (existingUser != models.User{}) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": loginUser.Email + " - Email already exists",
			})
		}

		// generate hash password
		hashedPassword, err := helpers.GenerateHashPassword(loginUser.Password)

		if err != nil {
			helpers.LogError(err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to Hash password",
			})
		}
		loginUser.Password = hashedPassword
		// insert into table
		loginUserID, err := models.AddUser(loginUser)
		if err != nil {
			helpers.LogError(err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to add a new user",
			})
		} else {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"message": "New User added successfully with id - " + strconv.Itoa(loginUserID),
			})
		}
	}
}

// Login Handler
/*
Get email and password from user
Validate email and password based on Login struct rules
get user details by provided email id if not found provide user not found message
if found check the provided password match with stored password(hashed version)
If password mismatch throw invalid credential message
Provide a JWT token for valid credentials
*/
func LoginHandler(c *fiber.Ctx) error {
	var loginUser Login
	// parse request body
	if err := c.BodyParser(&loginUser); err != nil {
		helpers.LogError(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Please provide valid inputs",
		})
	} else {
		// validate the user inputs
		errors := ValidateLoginUserStruct(loginUser)
		if errors != nil {
			return c.Status(fiber.StatusBadRequest).JSON(errors)
		}
		// Get user details by provided email id
		user, err := models.GetUserByEmail(loginUser.Email)
		if err != nil {
			helpers.LogError(err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to login a user",
			})
		} else {
			// If no user found, provide user not found message
			if (user == models.User{}) {
				return c.Status(fiber.StatusOK).JSON(fiber.Map{
					"message": "Login failed - user not found",
				})
			}
			// Check for password matching
			result := helpers.CheckHashPassword(loginUser.Password, user.Password)
			if result {
				// generate a JWT token
				token, _ := CreateJWT(loginUser.Email)
				return c.Status(fiber.StatusOK).JSON(fiber.Map{
					"message": "success",
					"token":   token,
				})
			} else {
				// throw invalid password message
				return c.Status(fiber.StatusOK).JSON(fiber.Map{
					"message": "Login failed - invalid password",
				})
			}
		}
	}
}

// Create JWT token
// Function to create JWT token
func CreateJWT(userEmail string) (string, error) {
	currentTime := time.Now().Format("02-01-2006 15:04:05")

	// Storing user name and loggedin time
	// Token expires in 1 hour.
	claims := MyCustomClaims{
		userEmail,
		currentTime,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(60 * time.Minute)),
			Issuer:    "Akilan",
		},
	}

	// Generate token with HS256 algorithm and custom claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Sign the token with our secret key
	signedToken, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))

	return signedToken, err
}

// Custom error message for invalid JWT
func JwtError(c *fiber.Ctx, err error) error {
	// Return status 401 and failed authentication error.
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "failed",
			"message": err.Error(),
		})
	}

	// Return status 401 and failed authentication error.
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"status":  "failed",
		"message": err.Error(),
	})
}
