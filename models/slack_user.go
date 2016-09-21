package models

import (
	"fmt"
)

// SlackUser stores slack user name and id
type SlackUser struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// NewSlackUser creates new SlackUser instance
func NewSlackUser(id string, name string) *SlackUser {
	return &SlackUser{
		ID:   id,
		Name: name,
	}
}

func (u SlackUser) String() string {
	return fmt.Sprintf("<@%v|%v>", u.ID, u.Name)
}
