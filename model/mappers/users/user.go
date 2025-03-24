package mapper

import (
	"go-api-template/model"
)

func ToUserModel(user *model.User) *model.User {
	return &model.User{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
