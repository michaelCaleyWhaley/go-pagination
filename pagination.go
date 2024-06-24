package main

import (
	"pagination/controllers"
	"pagination/initialisers"

	"github.com/gin-gonic/gin"
)

func init() {
	initialisers.LoadEnvVars()
	initialisers.ConnectToDB()
	initialisers.SyncDB()
	initialisers.CreatePeople()
}

func main() {

	app := gin.Default()

	app.LoadHTMLGlob("templates/**/*")
	app.Static("/assets", "./assets")

	app.GET("/", controllers.PeopleIndexGET)
	app.GET("/page/:page", controllers.PeopleIndexGET)

	app.Run()
}
