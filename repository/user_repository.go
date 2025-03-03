package repository

import "backend_architecture_golang/domains"

type UserRepository interface {
    GetAll() ([]domains.User, error)
    Create(user domains.User) error
}
