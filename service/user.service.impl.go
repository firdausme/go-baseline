package service

import (
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	Env "go-baseline/config/env"
	"go-baseline/constant"
	"go-baseline/entity"
	"go-baseline/exception"
	"go-baseline/model"
	"go-baseline/repository"
	"go-baseline/utils"
	"go-baseline/validation"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"time"
)

func NewUserService(userRepository *repository.UserRepository) UserService {
	return &userServiceImpl{
		UserRepository: *userRepository,
	}
}

type userServiceImpl struct {
	UserRepository repository.UserRepository
}

func (service userServiceImpl) Login(request model.LoginRequest) (response model.LoginResponse) {
	validation.LoginValidate(request)

	user, err := service.UserRepository.FindByUsername(request.Username)
	if err != nil {
		exception.PanicIfNeeded(err)
	}

	ok, err := checkPasswordHash(user.Password, request.Password)
	exception.PanicIfNeeded(err)

	if !ok {
		return
	}

	response = model.LoginResponse{
		Token: generateToken(user.Id, user.Username),
	}

	return response
}

func (service userServiceImpl) FindAll() (responses []model.GetUserResponse) {
	users := service.UserRepository.FindAll()
	for _, user := range users {
		responses = append(responses, model.GetUserResponse{
			Id:       user.Id,
			Username: user.Username,
			Password: user.Password,
		})
	}
	return responses
}

func (service userServiceImpl) Create(request model.CreateUserRequest) (response model.CreateUserResponse) {
	validation.CreateUserValidate(request)

	id := strings.ToUpper(uuid.New().String())

	user := entity.User{
		Id:       id,
		Username: request.Username,
		Password: request.Password,
	}

	service.UserRepository.Create(user)

	response = model.CreateUserResponse{
		Id:       id,
		Username: request.Username,
		Password: request.Password,
	}

	return response
}

func (service userServiceImpl) Update(request model.UpdateUserRequest) (response model.UpdateUserResponse) {
	validation.UpdateUserValidate(request)

	hashPwd := utils.HashPassword(request.Password)

	user := entity.User{
		Id:       request.Id,
		Username: request.Username,
		Password: hashPwd,
	}

	service.UserRepository.Update(user)

	response = model.UpdateUserResponse{
		Username: request.Username,
		Password: hashPwd,
	}

	return response
}

func (service userServiceImpl) Delete(request model.DeleteUserRequest) {
	validation.DeleteUserValidate(request)

	service.UserRepository.Delete(request.Id)
}

func generateToken(id, username string) string {
	newJwt := jwt.New(jwt.SigningMethodHS256)
	claims := newJwt.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	env := Env.Get()
	token, err := newJwt.SignedString([]byte(env[constant.JwtSecret]))
	exception.PanicIfNeeded(err)

	return token
}

func checkPasswordHash(hashPass, password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(password))
	exception.PanicIfNeeded(err)

	return true, nil
}
