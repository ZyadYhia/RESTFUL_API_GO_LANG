package routes

import (
	"example.com/rest-api/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	authenticated := server.Group("/")
	authenticated.Use(middleware.Authenticate)
	authenticated.POST("/events", createEvents)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/register", cancelForEvent)

	server.GET("/events", getEvents)
	server.GET("/event/:id", getEvent)
	//server.POST("/events", middleware.Authenticate, createEvents)
	//server.PUT("/event/:id", updateEvent)
	//server.DELETE("/event/:id", deleteEvent)
	server.POST("/signup", signup)
	server.POST("/login", login)
}
