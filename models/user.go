package models

import (
	"errors"
	"fmt"
)

// The User model type
type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User
	nextID = 1
)

// GetUsers function returns the users list
func GetUsers() []*User {
	return users
}

// AddUser function adds a new user to the list
func AddUser(user User) (User, error) {
	if user.ID != 0 {
		return User{}, errors.New("New User must not include id")
	}

	user.ID = nextID
	nextID++

	users = append(users, &user)

	return user, nil
}

// GetUserByID returns user from list by its id
func GetUserByID(id int) (User, error) {
	for _, user := range users {
		if user.ID == id {
			return *user, nil
		}
	}

	return User{}, fmt.Errorf("User with ID '%v' not found", id)
}

// UpdateUser updates an existent user from the list
func UpdateUser(user User) (User, error) {
	for i, cadidate := range users {
		if cadidate.ID == user.ID {
			users[i] = &user
			return user, nil
		}
	}

	return User{}, fmt.Errorf("User with ID '%v' not found", user.ID)
}

// RemoveUserByID removes user from the list
func RemoveUserByID(id int) error {
	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("User with ID '%v' not found", id)
}
