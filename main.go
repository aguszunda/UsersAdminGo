package main

import (
	controllers "agustinzunda/usersadmingo/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	pingController := controllers.NewUserController()

	r.GET("/ping", pingController.PingHandler)

	r.Run(":8080")
}
