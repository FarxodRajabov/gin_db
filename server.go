package main

import (
	"fmt"
	"gin_project/connection"
	"gin_project/controller"
	"gin_project/router"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	db := connection.Connection()

	handler := controller.NewHandler(db)

	fmt.Println("hasbeen started ...")
	router.Routers(server, handler)
	server.Run("localhost:2200")
	defer db.Close()
}
