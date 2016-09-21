package service

import (
	"github.com/wantedly/slack-mention-converter/models"
	"github.com/wantedly/slack-mention-converter/store"
)

// GetUser returns user by login name
func GetUser(s store.Store, loginName string) (*models.User, error) {
	return s.GetUser(loginName)
}

// AddUser adds or replaces user
func AddUser(s store.Store, user *models.User) error {
	return s.AddUser(user)
}

// ListUsers returns user list
func ListUsers(s store.Store) ([]*models.User, error) {
	return s.ListUsers()
}
