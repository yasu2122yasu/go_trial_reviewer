package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

type Information struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Pin     int    `json:"pin"`
}

var informations = []Information{
	{ID: "1", Name: "Satou", Address: "Tokyo", Pin: 7771414},
	{ID: "2", Name: "Netflix", Address: "Osaka", Pin: 5739111},
	{ID: "3", Name: "Microsoft", Address: "Kyoto", Pin: 6550013},
}

func Top(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"200": "The Index Method is passed"})
}

func Create(c *gin.Context) {
	var information Information
	if err := c.ShouldBindJSON(&information); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	information.ID = xid.New().String()
	informations = append(informations, information)
	c.JSON(http.StatusCreated, information)
}

func Index(c *gin.Context) {
	c.JSON(http.StatusOK, informations)
}

func Update(c *gin.Context) {
	id := c.Param("id")
	var information Information
	if err := c.ShouldBindJSON(&information); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	index := -1
	for i := 0; i < len(informations); i++ {
		if informations[i].ID == id {
			index = 1
		}
	}
	if index == -1 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Info not found",
		})
		return
	}
	informations[index] = information
	c.JSON(http.StatusOK, information)
}

func Delete(c *gin.Context) {
	id := c.Param("id")
	index := -1
	for i := 0; i < len(informations); i++ {
		if informations[i].ID == id {
			index = 1
		}
	}
	if index == -1 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Info not found",
		})
		return
	}
	informations = append(informations[:index], informations[index+1:]...)
	c.JSON(http.StatusOK, gin.H{
		"message": "Information has been deleted",
	})
}

func main() {
	router := gin.Default()
	router.GET("/", Top)
	router.GET("/informations", Index)
	router.POST("/information", Create)
	router.PUT("/information/:id", Update)
	router.DELETE("/information/:id", Delete)
	router.Run()
}
