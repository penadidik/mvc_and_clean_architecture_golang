package repository

import (
    "backend_architecture_golang/domains"
    "backend_architecture_golang/database"
)

type UserRepositoryImpl struct {}

func (u UserRepositoryImpl) GetAll() ([]domain.User, error) {
    var users []domain.User
    result := database.DB.Find(&users)
    return users, result.Error
}

func (u UserRepositoryImpl) Create(user domain.User) error {
    result := database.DB.Create(&user)
    return result.Error
}

