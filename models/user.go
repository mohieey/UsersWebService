package models

import (
	"errors"
	"fmt"
)

type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  = make(map[int]*User)
	nextID = 1
)

func GetUsers() map[int]*User {
	return users
}

func AddUser(u User) (User, error) {
	if u.ID != 0 {
		return User{}, errors.New("new user must not include an ID")
	}
	u.ID = nextID
	nextID++
	users[u.ID] = &u
	return u, nil
}

func GetUserByID(id int) (User, error) {
	user, exists := users[id]

	if !exists {
		return User{}, fmt.Errorf("User with ID %d not found", id)
	}

	return *user, nil
}

func DeleteUserByID(id int) error {
	_, exists := users[id]

	if !exists {
		return fmt.Errorf("User with ID %d not found", id)
	}

	delete(users, id)
	return nil
}

func UpdateUser(u User) (User, error) {
	_, exists := users[u.ID]

	if !exists {
		return User{}, fmt.Errorf("User with ID %d not found", u.ID)
	}

	users[u.ID] = &u

	return u, nil
}
