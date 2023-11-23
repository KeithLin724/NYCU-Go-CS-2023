package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Pages int    `json:"pages"`
}

type OnlyBook struct {
	Name  string `json:"name"`
	Pages int    `json:"pages"`
}

var bookshelf = []Book{
	{
		ID:    1,
		Name:  "Blue Bird",
		Pages: 500,
	},
}

func getBooks(c *gin.Context) {

	var bookshelfRe = []Book{}

	for _, item := range bookshelf {
		if item.Name == "delete" {
			continue
		}
		bookshelfRe = append(bookshelfRe, item)
	}

	c.JSON(http.StatusOK, bookshelfRe)
}

func getBook(c *gin.Context) {
	id := c.Param("id")

	idNumber, _ := strconv.ParseInt(id, 10, 64)

	length := int64(len(bookshelf))

	idNumber--

	if idNumber >= length || bookshelf[idNumber].Name == "delete" {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "book not found",
		})
		return
	}

	// replay
	c.JSON(http.StatusOK, bookshelf[idNumber])
}

func addBook(c *gin.Context) {

	var newBook OnlyBook

	// to json
	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// find have same
	for _, item := range bookshelf {
		if item.Name == newBook.Name && item.Pages == newBook.Pages {
			c.JSON(http.StatusConflict, gin.H{
				"message": "duplicate book name",
			})
			return
		}
	}

	// add book
	newBookIn := Book{
		ID:    len(bookshelf) + 1,
		Name:  newBook.Name,
		Pages: newBook.Pages,
	}

	bookshelf = append(bookshelf, newBookIn)

	c.JSON(http.StatusCreated, newBookIn)

}

func deleteBook(c *gin.Context) {
	id := c.Param("id")
	idNumber, _ := strconv.ParseInt(id, 10, 64)
	idNumber--

	length := int64(len(bookshelf))

	if idNumber >= length || bookshelf[idNumber].Name == "delete" {
		//return code 204
		c.Status(http.StatusNoContent)
		return
	}

	bookshelf[idNumber].Name = "delete"

	c.Status(http.StatusNoContent)

}

func updateBook(c *gin.Context) {

	id := c.Param("id")
	idNumber, _ := strconv.ParseInt(id, 10, 64)
	idNumber--

	var onlyBook OnlyBook
	// to json
	if err := c.ShouldBindJSON(&onlyBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if idNumber >= int64(len(bookshelf)) || bookshelf[idNumber].Name == "delete" {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "book not found",
		})
		return
	}

	oldBook := bookshelf[idNumber]

	// check have same name
	for _, item := range bookshelf {
		if item.Name == onlyBook.Name {
			c.JSON(http.StatusConflict, gin.H{
				"message": "duplicate book name",
			})
			return
		}
	}

	// same book
	if oldBook.Name == onlyBook.Name && oldBook.Pages == onlyBook.Pages {
		c.JSON(http.StatusConflict, gin.H{
			"message": "duplicate book name",
		})
		return
	}

	bookshelf[idNumber].Name = onlyBook.Name
	bookshelf[idNumber].Pages = onlyBook.Pages
	// return
	c.JSON(http.StatusOK, bookshelf[idNumber])
}

func main() {
	r := gin.Default()
	r.RedirectFixedPath = true

	// TODO: Add routes
	r.GET("/bookshelf", getBooks)
	r.GET("/bookshelf/:id", getBook)
	r.POST("/bookshelf", addBook)
	r.PUT("/bookshelf/:id", updateBook)
	r.DELETE("/bookshelf/:id", deleteBook)

	err := r.Run(":8087")
	if err != nil {
		return
	}
}
