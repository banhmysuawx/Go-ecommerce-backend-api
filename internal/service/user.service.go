package service

import "github.com/banhmysuawx/Go-ecommerce-backend-api/internal/repo"

type UserService struct {
	userRepo *repo.UserRepo
}

func NewUserService() *UserService {
	return &UserService{
		userRepo: repo.NewUserRepo(),
	}
}

func (us *UserService) GetInfoUserByID(id string) string {
	return us.userRepo.GetInfoUserByID(id)
}