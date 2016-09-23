package service

import (
	"github.com/wantedly/slack-mention-converter/models"
	"github.com/wantedly/slack-mention-converter/store"
)

// GetSlackUser returns slack user by its name
func GetSlackUser(s store.Store, name string) (*models.SlackUser, error) {
	return s.GetSlackUser(name)
}

// ListSlackUsers returns slack user list
func ListSlackUsers(s store.Store) ([]*models.SlackUser, error) {
	return s.ListSlackUsers()
}
