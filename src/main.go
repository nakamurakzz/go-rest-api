package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type item struct {
	ID         string  `json:"id"`
	Title      string  `json:"title"`
	Price      float64 `json:"price"`
	ItemTypeId int     `json:"item_type_id"`
}

var items = []item{
	{ID: "1", Title: "Item 1", Price: 1.99, ItemTypeId: 1},
	{ID: "2", Title: "Item 2", Price: 2.99, ItemTypeId: 1},
	{ID: "3", Title: "Item 3", Price: 3.99, ItemTypeId: 2},
}

func getItems(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, items)
}

func main() {
	router := gin.Default()
	router.GET("/items", getItems)

	router.Run(":8080")
}
