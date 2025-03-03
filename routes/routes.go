package routes

import (
    "backend_architecture_golang/controllers"
    "backend_architecture_golang/repository"
    "backend_architecture_golang/usecase"
    "github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()

    userRepo := repository.UserRepositoryImpl{}
    userUC := usecase.UserUsecase{userRepo}
    userController := controllers.UserController{userUC}

    r.GET("/users", userController.GetUsers)
    r.POST("/users", userController.CreateUser)

    return r
}

