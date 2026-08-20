package main 

import "errors"

type UserService struct {
	repository	UserRepository
}


func NewUserService(repository UserRepository) *UserService {
	return &UserService{
		repository: repository,
	}
}

func (s *UserService) Register(
	username string,
	email string,
	password string,
) (*User, error) {

	if username == "" || email == "" || password == "" {
		return "", ErrInvalidCredentials
	}

	hash, err := HashPassword(password)

	if err != nil {
		return nil, err
	}

	user := &User {
		Username: username,
		Email:	email,
		Password: hash,
	}

	return s.repository.Create(user)
}

func (s *UserService) Authenticate(email string,password string) (*User, error) {
	user, ok := s.repository.GetByEmail(email)

	if !ok {
		return nil, ErrInvalidCredentials
	}

	if err:= CheckPassword(password); err!=nil{
		return nil,ErrInvalidCredentials
	}
	return user,nil
}


