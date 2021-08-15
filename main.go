package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// book represents data about a single book.
type book struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
}

// books slice to seed book data. In memory database
var books = []book{
	{ID: "1", Title: "The Lord of the Rings - The fellowship of the ring", Author: "J. R. R. Tolkien", Price: 1.00},
	{ID: "2", Title: "The Lord of the Rings  - The two towers", Author: "J. R. R. Tolkien", Price: 2.00},
	{ID: "3", Title: "The Lord of the Rings  - The return of the king", Author: "J. R. R. Tolkien", Price: 3.00},
}

// getBooks responds with the list of all books as JSON.
func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

// postBooks adds an book from JSON received in the request body.
func postBooks(c *gin.Context) {
	var newBook book

	// Call BindJSON to bind the received JSON to newBook.
	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	// Add the new book to the slice.
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

// getBookByID locates the book whose ID value matches the id
// parameter sent by the client, then returns that book as a response.
func getBookByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of books, looking for
	// an book whose ID value matches the parameter.
	for _, a := range books {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id", getBookByID)
	router.POST("/books", postBooks)

	router.Run("localhost:8080")
}