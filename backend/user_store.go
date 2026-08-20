package main

import (
	"strings"
	"sync"
)

type UserStore struct {
	mu		sync.RWMutex
	users 	map[int]*User
	nextID	int
}

func NewUserStore() *UserStore {
	return &UserStore {
		users: make(map[int]*User),
		nextID: 1,
	}
}

func (s *UserStore) Create(user *User) (*User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _,existing := range s.users {
		if strings.EqualFold(existing.Email,user.Email) {
			return nil, ErrEmailAlreadyExists
		}
	}
	user.ID = s.nextID
	s.nextID++
	s.users[user.ID] = user

	return user,nil
}

func (s *UserStore) GetByID(id int) (*User, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	user, ok := s.users[id]


	return user, ok
}

func (s *UserStore)GetByEmail(email string) (*User, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, user := range s.users {
		if strings.EqualFold(user.Email,email) {
			return user, true
		}
	}
	return nil, false
}



