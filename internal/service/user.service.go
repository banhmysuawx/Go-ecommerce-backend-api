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

func (us *UserService) GetInfoUserByID(id string) (string, error) {
	result := us.userRepo.GetInfoUserByID(id)
	if result == "" { // Assuming an empty string indicates no user found or an error
		return "", nil // or return "", errors.New("user not found") to indicate an error
	}
	return result, nil
}