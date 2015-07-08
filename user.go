package main

import (
	"fmt"
)

// User represents a Slack user
type User struct {
	UserName   string
	UserID     string
	OAuthToken string
}

// UserStore is a data structure to store users
type UserStore map[string]*User

// Add a user to the store
func (store *UserStore) Add(user *User) error {
	if len(user.UserID) == 0 {
		return fmt.Errorf("Invalid UserID")
	}
	_, exists := (*store)[user.UserID]
	if exists {
		return fmt.Errorf("User already exists")
	}

	(*store)[user.UserID] = user

	return nil
}

// GetOrAdd returns a user if it exists or creates a new one.
// The second return parameter is TRUE if user has been created
func (store *UserStore) GetOrAdd(userID string) (*User, bool) {
	user, exists := (*store)[userID]
	if exists {
		return user, false
	}
	user = &User{UserID: userID}

	return user, true
}
