package service

import (
	"encoding/csv"
	"errors"
	"io"
	"os"
	"path/filepath"

	"github.com/wantedly/slack-mention-converter/models"
)

var (
	// UserMapCachePath is the file path to store login name and slack name pairs as csv
	UserMapCachePath = "data/user_map.csv"
)

func cacheFileUserMapPath() string {
	curDir, _ := os.Getwd()
	return filepath.Join(curDir, UserMapCachePath)
}

// GetUser returns user by login name
func GetUser(loginName string) (*models.User, error) {
	users, err := ListUsers()
	if err != nil {
		return &models.User{}, err
	}
	for _, user := range users {
		if user.LoginName == loginName {
			return user, nil
		}
	}
	return &models.User{}, errors.New("Such login name not found")
}

// AddUser adds or replaces user
func AddUser(user *models.User) error {
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
func ListUsers() ([]*models.User, error) {
	file, err := os.Open(cacheFileUserMapPath())
	if err != nil {
		return []*models.User{}, nil
	}
	defer file.Close()
	reader := csv.NewReader(file)

	var res []*models.User
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return res, err
		}
		res = append(res, models.NewUser(record[0], record[1]))
	}
	return res, nil
}

func putUsers(users []*models.User) error {
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
