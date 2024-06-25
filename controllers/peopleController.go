package controllers

import (
	"net/http"
	"pagination/helpers"
	"pagination/initialisers"
	"pagination/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PaginationData struct {
	NextPage        int
	PreviousPage    int
	CurrentPage     int
	TotalPages      int
	TwoPreviousPage int
	TwoNextPage     int
	ThreeNextPage   int
}

func PeopleIndexGET(c *gin.Context) {
	pageNo := 1
	perPage := 10
	pageString := c.Param("page")

	if pageString != "" {
		pageNo, _ = strconv.Atoi(pageString)
	}

	pagination := helpers.GetPagintionData(pageNo, perPage, &models.Person{}, "/people")

	var people []models.Person
	initialisers.DB.Limit(perPage).Offset(pagination.Offset).Find(&people)

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"people":     people,
		"pagination": pagination,
	})
}
