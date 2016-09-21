package store

import (
	"github.com/wantedly/slack-mention-converter/models"
)

type Store interface {
	GetUser(loginName string) (*models.User, error)
	AddUser(user *models.User) error
	ListUsers() ([]*models.User, error)
	PutUsers(users []*models.User) error
	GetSlackUser(name string) (*models.SlackUser, error)
	ListSlackUsers() ([]*models.SlackUser, error)
}
