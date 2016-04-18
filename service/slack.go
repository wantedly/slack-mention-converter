package service

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/jmoiron/jsonq"
)

const (
	// SlackUserListURL represents users.list Slack API endpoint (https://api.slack.com/methods/users.list)
	SlackUserListURL = "https://slack.com/api/users.list"
)

// SlackUser stores slack user name and id
type SlackUser struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// NewSlackUser creates new SlackUser instance
func NewSlackUser(id string, name string) SlackUser {
	return SlackUser{
		ID:   id,
		Name: name,
	}
}

// ListSlackUsers returns slack user list
func ListSlackUsers() []SlackUser {
	requestURL := SlackUserListURL + "?token=" // TODO(awakia) add slack token
	resp, err := http.Get(requestURL)
	if err != nil {
		log.Println(err)
		return []SlackUser{}
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
	var res []SlackUser
	for i := 0; i < len(arr); i++ {
		id, _ := jq.String("members", strconv.Itoa(i), "id")
		name, _ := jq.String("members", strconv.Itoa(i), "name")
		res = append(res, NewSlackUser(id, name))
	}

	return res
}
