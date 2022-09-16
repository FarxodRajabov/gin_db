package router

import (
	"gin_project/controller"

	"github.com/gin-gonic/gin"
)

func Routers(c *gin.Engine, h controller.UserHandler) {
	c.GET("/users", h.GetUsers)
	c.POST("/users", h.PostUser)
	c.GET("/users/:id", h.GetByID)
	c.DELETE("/users/:id", h.DeleteById)
	c.PUT("/users", h.UpdateUser)
}
