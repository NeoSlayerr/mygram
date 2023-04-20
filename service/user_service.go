package service

import (
	"fmt"
	"mygram/dto"
	"mygram/entity"
	"mygram/pkg/errs"
	"mygram/pkg/helpers"
	"mygram/repository/user_repository"
	"net/http"
)

type userService struct {
	userRepo user_repository.UserRepository
}

type UserService interface {
	CreateNewUser(payload dto.NewUserRequest) (*dto.NewUserResponse, errs.MessageErr)
	Login(payload dto.NewLoginRequest) (*dto.LoginResponse, errs.MessageErr)
}

func NewUserService(userRepo user_repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (u *userService) Login(payload dto.NewLoginRequest) (*dto.LoginResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(payload)

	if err != nil {
		return nil, err
	}

	user, err := u.userRepo.GetUserByEmail(payload.Email)

	if err != nil {
		if err.Status() == http.StatusNotFound {
			return nil, errs.NewUnauthenticatedError("invalid email/password")
		}
		return nil, err
	}

	isValidPassword := user.ComparePassword(payload.Password)

	if !isValidPassword {
		return nil, errs.NewUnauthenticatedError("invalid email/password")
	}

	fmt.Println("user =>", user)

	response := dto.LoginResponse{
		Result:     "success",
		Message:    "logged in successfully",
		StatusCode: http.StatusOK,
		Data: dto.TokenResponse{
			Token: user.GenerateToken(),
		},
	}

	return &response, nil
}

func (u *userService) CreateNewUser(newUserRequest dto.NewUserRequest) (*dto.NewUserResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(newUserRequest)

	if err != nil {
		return nil, err
	}


	payload := entity.User{
		Username:    newUserRequest.Username,
		Email:    newUserRequest.Email,
		Password: newUserRequest.Password,
		Age: newUserRequest.Age,
	}

	err = payload.HashPassword()

	if err != nil {
		return nil, err
	}

	err = u.userRepo.CreateNewUser(payload)

	if err != nil {
		return nil, err
	}

	response := dto.NewUserResponse{
		Result:     "success",
		StatusCode: http.StatusCreated,
		Message:    "registered successfully",
	}

	return &response, nil
}
