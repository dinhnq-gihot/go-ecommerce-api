package services

import "go-ecommerce-backend-api/internal/repos"

type UserService struct {
	repo *repos.UserRepo
}

func NewUserService() *UserService {
	return &UserService{
		repo: repos.NewUserRepo(),
	}
}

func (us *UserService) GetUserById(id int) string {
	return us.repo.GetUserById(id)
}
