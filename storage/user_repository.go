package storage

import (
	"github.com/Stremilov/lk-techdep-backend/internal/user"
)

type UserRepository struct{}

func (r *UserRepository) CreateUser(u user.User) error {
	// Логика для добавления пользователя в БД
}

func (r *UserRepository) GetUserByID(id string) (user.User, error) {
	// Логика для получения пользователя из БД
}

// Другие методы...
