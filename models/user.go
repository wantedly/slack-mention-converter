package models

import (
	"fmt"
)

// User stores login user name and slack user name
type User struct {
	LoginName string
	SlackName string
}

// NewUser creates new UserMap instance
func NewUser(loginName string, slackName string) *User {
	return &User{
		LoginName: loginName,
		SlackName: slackName,
	}
}

func (u *User) String() string {
	return fmt.Sprintf("%v:@%v", u.LoginName, u.SlackName)
}
