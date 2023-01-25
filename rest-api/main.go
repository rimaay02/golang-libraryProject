package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"

)
func postBooks(c *gin.Context) {
	var newBook Book

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newBook); err != nil {
			return
	}

	// Add the new album to the slice.
	Books = append(Books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}
type Book struct {
	ID     string  `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
}
func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Books)
}
var Books = []Book{
	{ID: "1", Title: "Harry Potter", Author: "JK Rowling"},
	{ID: "2",Title: "Shopaholic", Author: "Sophie Kinsella"},	
}

func main() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
        AllowOrigins: []string{"*"},
        AllowMethods: []string{"POST", "PUT", "PATCH", "DELETE"},
        AllowHeaders: []string{"Content-Type,access-control-allow-origin, access-control-allow-headers"},
    }))
	router.GET("/",getBooks )
	router.POST("/books", postBooks)
	router.Run("localhost:8081")
}