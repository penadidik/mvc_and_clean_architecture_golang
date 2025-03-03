package routes

import (
    "backend_architecture_golang/controllers"
    "github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()
    r.GET("/users", controllers.GetUsers)
    r.POST("/users", controllers.CreateUser)
    return r
}
