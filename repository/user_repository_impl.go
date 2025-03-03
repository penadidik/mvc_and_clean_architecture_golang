package repository

import (
    "backend_architecture_golang/domains"
    "backend_architecture_golang/database"
)

type UserRepositoryImpl struct {}

func (u UserRepositoryImpl) GetAll() ([]domains.User, error) {
    var users []domains.User
    result := database.DB.Find(&users)
    return users, result.Error
}

func (u UserRepositoryImpl) Create(user domains.User) error {
    result := database.DB.Create(&user)
    return result.Error
}

