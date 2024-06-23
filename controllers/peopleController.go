package controllers

import (
	"net/http"
	"pagination/initialisers"
	"pagination/models"

	"github.com/gin-gonic/gin"
)

func PeopleIndexGET(c *gin.Context) {
	// Get the people
	var people []models.Person
	initialisers.DB.Find(&people)

	// Render the page
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"people": people,
	})
}
