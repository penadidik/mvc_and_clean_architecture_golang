package repository

import "backend_architecture_golang/domains"

type UserRepository interface {
    GetAll() ([]domain.User, error)
    Create(user domain.User) error
}

