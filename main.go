package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type item struct {
	Id         string  `json:"id"`
	Title      string  `json:"title"`
	Price      float64 `json:"price"`
	ItemTypeId int     `json:"item_type_id"`
}

func newItem(title string, price float64, itemTypeId int) (item, error) {
	// idをuuidで生成
	id, err := uuid.NewRandom()
	if err != nil {
		return item{}, err
	}

	return item{
		Id:         id.String(),
		Title:      title,
		Price:      price,
		ItemTypeId: itemTypeId,
	}, nil
}

type itemReq struct {
	Title      string  `json:"title"`
	Price      float64 `json:"price"`
	ItemTypeId int     `json:"item_type_id"`
}

var items = []item{
	{Id: "5bdf0ad1-2d44-4794-b7de-a0ed40c81ab8", Title: "Item 1", Price: 1.99, ItemTypeId: 1},
	{Id: "c513d2eb-3deb-45d0-85a1-0be41f80f80c", Title: "Item 2", Price: 2.99, ItemTypeId: 1},
	{Id: "29910786-3d46-474e-b4a6-db4f600db75b", Title: "Item 3", Price: 3.99, ItemTypeId: 2},
}

func getItems(c *gin.Context) {
	c.JSON(http.StatusOK, items) // 200
	// c.JSON(http.StatusNotFound, items)            // 404
	// c.JSON(http.StatusInternalServerError, items) // 500
	// c.IndentedJSON(http.StatusOK, items) // 200, IndentされたJSON
}

func postItems(c *gin.Context) {
	var newItemReq itemReq
	if err := c.BindJSON(&newItemReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newItemEntity, err := newItem(newItemReq.Title, newItemReq.Price, newItemReq.ItemTypeId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	items = append(items, newItemEntity)

	c.JSON(http.StatusCreated, items) // 201
}

func putItem(c *gin.Context) {
	id := c.Param("id")

	var newItem itemReq
	if err := c.BindJSON(&newItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, item := range items {
		if item.Id == id {
			items[i].Title = newItem.Title
			items[i].Price = newItem.Price
			items[i].ItemTypeId = newItem.ItemTypeId
			c.JSON(http.StatusOK, items[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "item not found"})
}

func main() {
	router := gin.Default()
	router.GET("/items", getItems)
	router.POST("/items", postItems)
	router.PUT("/items:id", putItem)

	router.Run(":8080")
}
