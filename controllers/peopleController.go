package controllers

import (
	"net/http"
	"pagination/initialisers"
	"pagination/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PaginationData struct {
	NextPage     int
	PreviousPage int
	CurrentPage  int
	TotalPages   int
}

func PeopleIndexGET(c *gin.Context) {
	pageNo := 1
	perPage := 10
	pageString := c.Param("page")

	if pageString != "" {
		pageNo, _ = strconv.Atoi(pageString)
	}

	var totalRows int64
	initialisers.DB.Model(&models.Person{}).Count(&totalRows)
	totalPages := float64(totalRows / int64(perPage))

	offset := (pageNo - 1) * perPage

	var people []models.Person
	initialisers.DB.Limit(perPage).Offset(offset).Find(&people)

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"people":     people,
		"pagination": PaginationData{NextPage: pageNo + 1, PreviousPage: pageNo - 1, CurrentPage: pageNo, TotalPages: int(totalPages)},
	})
}
