package main

import (
	"agustinzunda/usersadmingo/internal/db"
)

func main() {
	db.ConnDb()
}

/*func main() {
	r := gin.Default()

	pingController := controllers.NewUserController()

	r.GET("/ping", pingController.PingHandler)

	r.Run(":8080")

}*/
