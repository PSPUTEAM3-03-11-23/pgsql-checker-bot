package service

import (
	"hackathon-tg-bot/internal/app/mapper"
	"hackathon-tg-bot/internal/app/model/entity"
	"hackathon-tg-bot/internal/app/model/input"
	"hackathon-tg-bot/internal/app/model/response"
	"hackathon-tg-bot/pkg/errorHandler"
	"net/http"
)

type UserRepository interface {
	GetByCreds(loginInput *input.Login) (*entity.User, error)
}
type AuthService struct {
	userRepo UserRepository
}

func NewAuthService(userRepo UserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

func (d *AuthService) Login(loginInput *input.Login) (*response.User, *errorHandler.HttpErr) {
	userResponse := &response.User{}

	user, err := d.userRepo.GetByCreds(loginInput)
	if user == nil {
		return nil, errorHandler.New("User does not exists", http.StatusForbidden)
	}
	if err != nil {
		return nil, errorHandler.New(err.Error(), http.StatusBadRequest)
	}

	userResponse = mapper.UserToUserResponse(user)

	return userResponse, nil
}
