package controllers

import (
    "backend_architecture_golang/domains"
    "backend_architecture_golang/usecase"
    "github.com/gin-gonic/gin"
    "net/http"
)

type UserController struct {
    userUsecase usecase.UserUsecase
}

func (uc UserController) GetUsers(c *gin.Context) {
    users, err := uc.userUsecase.GetAllUsers()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, users)
}

func (uc UserController) CreateUser(c *gin.Context) {
    var user domains.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    err := uc.userUsecase.CreateUser(user)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, user)
}

