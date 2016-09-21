package store

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/jmoiron/jsonq"
	"github.com/wantedly/slack-mention-converter/models"
)

const (
	slackUserListURL = "https://slack.com/api/users.list"
	slackIDsName     = "slack_users.csv"
	userMapName      = "user_map.csv"
)

type CSV struct {
	dir string
}

func NewCSV(dir string) *CSV {
	return &CSV{
		dir: dir,
	}
}

func (c *CSV) userMapPath() string {
	return filepath.Join(c.dir, userMapName)
}

func (c *CSV) slackUsersPath() string {
	return filepath.Join(c.dir, slackIDsName)
}

func (c *CSV) GetUser(loginName string) (*models.User, error) {
	users, err := c.ListUsers()
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

func (c *CSV) AddUser(user *models.User) error {
	users, _ := c.ListUsers()
	for i := 0; i < len(users); i++ {
		if users[i].LoginName == user.LoginName {
			users[i] = user
			return c.PutUsers(users)
		}
	}
	users = append(users, user)
	return c.PutUsers(users)
}

func (c *CSV) ListUsers() ([]*models.User, error) {
	file, err := os.Open(c.userMapPath())
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

func (c *CSV) PutUsers(users []*models.User) error {
	file, err := os.OpenFile(c.userMapPath(), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
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

func (c *CSV) GetSlackUser(name string) (*models.SlackUser, error) {
	slackUsers, err := c.getSlackUsersFromCache()
	if err != nil {
		return &models.SlackUser{}, err
	}
	for _, user := range slackUsers {
		if user.Name == name {
			return user, nil
		}
	}

	slackUsers, err = c.fetchSlackUsers()
	for _, user := range slackUsers {
		if user.Name == name {
			return user, nil
		}
	}
	return &models.SlackUser{}, errors.New("Slack id not found")
}

func (c *CSV) ListSlackUsers() ([]*models.SlackUser, error) {
	cached, err := c.getSlackUsersFromCache()
	if len(cached) > 0 {
		return cached, err
	}
	return c.fetchSlackUsers()
}

func (c *CSV) getSlackUsersFromCache() ([]*models.SlackUser, error) {
	file, err := os.Open(c.slackUsersPath())
	if err != nil {
		return []*models.SlackUser{}, err
	}
	defer file.Close()
	reader := csv.NewReader(file)

	var res []*models.SlackUser
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return res, err
		}
		res = append(res, models.NewSlackUser(record[0], record[1]))
	}
	return res, nil
}

func (c *CSV) putSlackUsersToCache(slackUsers []*models.SlackUser) error {
	file, err := os.OpenFile(c.slackUsersPath(), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	writer := csv.NewWriter(file)

	for _, user := range slackUsers {
		writer.Write([]string{user.ID, user.Name})
	}
	writer.Flush()
	return nil
}

func (c *CSV) fetchSlackUsers() ([]*models.SlackUser, error) {
	token := os.Getenv("SLACK_API_TOKEN")
	if token == "" {
		log.Fatalf("You need to pass SLACK_API_TOKEN as environment variable.")
	}
	requestURL := slackUserListURL + "?token=" + token
	resp, err := http.Get(requestURL)
	if err != nil {
		return []*models.SlackUser{}, err
	}
	defer resp.Body.Close()

	data := map[string]interface{}{}
	dec := json.NewDecoder(resp.Body)
	dec.Decode(&data)
	jq := jsonq.NewQuery(data)
	arr, err := jq.Array("members")
	if err != nil {
		log.Println(err)
	}
	var res []*models.SlackUser
	for i := 0; i < len(arr); i++ {
		id, _ := jq.String("members", strconv.Itoa(i), "id")
		name, _ := jq.String("members", strconv.Itoa(i), "name")
		res = append(res, models.NewSlackUser(id, name))
	}

	putErr := c.putSlackUsersToCache(res)
	if putErr != nil {
		log.Println(putErr)
	}
	return res, nil
}
