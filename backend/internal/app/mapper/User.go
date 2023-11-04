package mapper

import (
	"hackathon-tg-bot/internal/app/model/entity"
	"hackathon-tg-bot/internal/app/model/response"
)

func UserToUserResponse(user *entity.User) *response.User {
	r := &response.User{
		Id:            user.Id,
		Email:         user.Email,
		Password:      user.Password,
		IsDeactivated: user.IsDeactivated,
	}

	return r
}

func UserToUserResponses(users *[]entity.User) *[]response.User {
	rs := make([]response.User, 0)

	for _, user := range *users {
		rs = append(rs, *UserToUserResponse(&user))
	}

	return &rs
}
