package models

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jmoiron/jsonq"
)

const (
	slackUserListURL = "https://slack.com/api/users.list"
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

// RetrieveFromSlack retrieves users via Slack API
func RetrieveFromSlack(token string) ([]*SlackUser, error) {
	if token == "" {
		return nil, fmt.Errorf("You need to pass SLACK_API_TOKEN as environment variable")
	}

	requestURL := slackUserListURL + "?token=" + token
	resp, err := http.Get(requestURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data := map[string]interface{}{}
	dec := json.NewDecoder(resp.Body)
	dec.Decode(&data)
	jq := jsonq.NewQuery(data)
	arr, err := jq.Array("members")
	if err != nil {
		return nil, err
	}

	var users []*SlackUser

	for i := 0; i < len(arr); i++ {
		id, _ := jq.String("members", strconv.Itoa(i), "id")
		name, _ := jq.String("members", strconv.Itoa(i), "name")
		users = append(users, NewSlackUser(id, name))
	}

	return users, nil
}

func (u *SlackUser) String() string {
	return fmt.Sprintf("<@%v|%v>", u.ID, u.Name)
}
