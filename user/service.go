package user

import (
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// interface
type IUserService interface {
	Save(input UserInputLogin) (User, error)
}

type UserService struct {
	userRepo IUserRepository
}

// new service
func NewService(userRepo IUserRepository) *UserService {
	return &UserService{userRepo}
}

// implementasi
func (s *UserService) Save(input UserInputLogin) (User, error) {
	// cari user by email
	userByEmail, err := s.userRepo.FindByEmail(input.Email)
	if err != nil {
		return userByEmail, err
	}

	// jangan ada email yang sama
	if userByEmail.Email == input.Email {
		errMsg := fmt.Sprintf("user dengan email %v telah terdaftar", input.Email)
		return userByEmail, errors.New(errMsg)
	}

	// binding
	var user User
	user.Email = input.Email
	user.Username = input.Username

	bytePass, err := bcrypt.GenerateFromPassword([]byte(input.Password), 10)
	if err != nil {
		errMsg := fmt.Sprintf("gagal enkripsi password : %v", err)
		return userByEmail, errors.New(errMsg)
	}

	user.Password = string(bytePass)

	// save
	userSaved, err := s.userRepo.Save(user)
	if err != nil {
		return userSaved, err
	}

	return userSaved, nil
}
