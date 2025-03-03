package usecase

import (
    "backend_architecture_golang/domains"
    "backend_architecture_golang/repository"
)

type UserUsecase struct {
    userRepo repository.UserRepository
}

func (u UserUsecase) GetAllUsers() ([]domain.User, error) {
    return u.userRepo.GetAll()
}

func (u UserUsecase) CreateUser(user domain.User) error {
    return u.userRepo.Create(user)
}

