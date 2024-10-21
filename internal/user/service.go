package user

import "github.com/Stremilov/lk-techdep-backend/storage"

type UserService struct {
	repo *storage.UserRepository
}

func NewUserService(repo *storage.UserRepository) *UserService {
	return &UserService{repo: repo}
}
