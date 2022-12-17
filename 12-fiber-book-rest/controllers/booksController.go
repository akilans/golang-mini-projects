package controllers

import (
	"strconv"

	"github.com/akilans/fiber-book-rest/helpers"
	"github.com/akilans/fiber-book-rest/models"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// Payload validation response message
type ErrorResponse struct {
	FailedField string `json:"failed_field"`
	Tag         string `json:"tag"`
	Value       string `json:"value"`
}

var validate = validator.New()

// validate book payload

func ValidateBookStruct(book models.Book) []ErrorResponse {
	var errors []ErrorResponse
	err := validate.Struct(book)
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

// List books function
func ListBooksHandler(c *fiber.Ctx) error {
	// Get list of books from DB
	books, err := models.GetBooks()
	if err != nil {
		helpers.LogError(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to list books",
		})
	} else {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": books,
		})
	}
}

// Add a new book function
/*
Get new book inputs from user
validate it
Store in the DB
*/
func AddBookHandler(c *fiber.Ctx) error {
	var newBook models.Book

	// Parse request body
	if err := c.BodyParser(&newBook); err != nil {
		helpers.LogError(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Please provide valid inputs",
		})
	} else {
		// Validate user inputs
		errors := ValidateBookStruct(newBook)
		if errors != nil {
			return c.Status(fiber.StatusBadRequest).JSON(errors)
		}
		// Add new book into db
		newBookID, err := models.AddBook(newBook)
		if err != nil {
			helpers.LogError(err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to add a new book",
			})
		} else {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"message": "New book added successfully with id - " + strconv.Itoa(newBookID),
			})
		}
	}
}

// Get book by id function
/*
Validate provided book id in the URL
Get book details from DB
If no book returned send book not found message
else give book info
*/
func GetBookHandler(c *fiber.Ctx) error {
	// Validate provided book id
	bookId, err := c.ParamsInt("id")
	if err != nil {
		helpers.LogError(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Please provide a valid book id",
		})
	}
	// Get book by id
	book := models.GetBookByID(bookId)

	// Check for empty result
	if (book == models.Book{}) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Book doesn't exists",
		})
	} else {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": book,
		})
	}

}

// Update a book by id function
/*
Validate provided book id in the URL
Get book details from DB
If no book returned send book not found message
else update book details
*/
func UpdateBookHandler(c *fiber.Ctx) error {

	// Validate provided book id
	bookId, err := c.ParamsInt("id")

	if err != nil {
		helpers.LogError(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Please provide a valid book id",
		})
	}

	// Get book by id
	book := models.GetBookByID(bookId)

	// Check for empty result
	if (book == models.Book{}) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Book doesn't exists",
		})
	}

	// parse the request body
	if err := c.BodyParser(&book); err != nil {
		helpers.LogError(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to update a book",
		})
	} else {
		// Update a book with new values
		updatedBookID, err := models.UpdateBook(book)
		if err != nil {
			helpers.LogError(err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to update a book",
			})
		} else {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"message": "Book updated successfully with id - " + strconv.Itoa(updatedBookID),
			})
		}
	}
}

// Delete books function
/*
Validate provided book id in the URL
If book doesn't exists send book not found message
else delete a book
*/
func DeleteBookHandler(c *fiber.Ctx) error {
	// Validate provided book id
	bookId, err := c.ParamsInt("id")
	if err != nil {
		helpers.LogError(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Please provide a valid book id",
		})
	}

	// Delete book by id
	err = models.DeleteBookByID(bookId)
	if err != nil {
		helpers.LogError(err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Book doesn't exists",
		})
	} else {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Book deleted successfully",
		})
	}
}
