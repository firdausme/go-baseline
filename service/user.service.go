package service

import "go-baseline/model"

type UserService interface {
	Login(request model.LoginRequest) (response model.LoginResponse)
	FindAll() (responses []model.GetUserResponse)
	Create(request model.CreateUserRequest) (response model.CreateUserResponse)
	Update(request model.UpdateUserRequest) (response model.UpdateUserResponse)
	Delete(request model.DeleteUserRequest)
}
