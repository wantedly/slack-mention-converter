package service

import (
	"encoding/csv"
	"errors"
	"io"
	"os"
	"path/filepath"
)

const (
	// UserMapCachePath is the file path to store login name and slack name pairs as csv
	UserMapCachePath = "data/user_map.csv"
)

// User stores login user name and slack user name
type User struct {
	LoginName string
	SlackName string
}

func cacheFileUserMapPath() string {
	curDir, _ := os.Getwd()
	return filepath.Join(curDir, "..", UserMapCachePath)
}

// NewUser creates new UserMap instance
func NewUser(loginName string, slackName string) User {
	return User{
		LoginName: loginName,
		SlackName: slackName,
	}
}

// GetUser returns user by login name
func GetUser(loginName string) (User, error) {
	users, err := ListUsers()
	if err != nil {
		return User{}, err
	}
	for _, user := range users {
		if user.LoginName == loginName {
			return user, nil
		}
	}
	return User{}, errors.New("Such login name not found")
}

// AddUser adds or replaces user
func AddUser(user User) error {
	users, _ := ListUsers()
	for i := 0; i < len(users); i++ {
		if users[i].LoginName == user.LoginName {
			users[i] = user
			return putUsers(users)
		}
	}
	users = append(users, user)
	return putUsers(users)
}

// ListUsers returns user list
func ListUsers() ([]User, error) {
	file, err := os.Open(cacheFileUserMapPath())
	if err != nil {
		return []User{}, err
	}
	defer file.Close()
	reader := csv.NewReader(file)

	var res []User
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return res, err
		}
		res = append(res, NewUser(record[0], record[1]))
	}
	return res, nil
}

func putUsers(users []User) error {
	file, err := os.OpenFile(cacheFileUserMapPath(), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	writer := csv.NewWriter(file)

	for _, user := range users {
		writer.Write([]string{user.LoginName, user.SlackName})
	}
	writer.Flush()
	return nil
}
