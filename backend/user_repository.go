package main

type UserRepository interface {
	Create(user *User) (*User, error)
	GetByID(id int) (*User, bool)
	GetByEmail(email string) (*User, bool)
}


