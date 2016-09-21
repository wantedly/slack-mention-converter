package service

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
	// SlackUserListURL represents users.list Slack API endpoint (https://api.slack.com/methods/users.list)
	SlackUserListURL = "https://slack.com/api/users.list"
	// SlackUserCachePath is the file path to store slack users ids and names as csv
	SlackUserCachePath = "data/slack_users.csv"
)

func cacheSlackUserFilePath() string {
	curDir, _ := os.Getwd()
	return filepath.Join(curDir, SlackUserCachePath)
}

// GetSlackUser returns slack user by its name
func GetSlackUser(name string) (*models.SlackUser, error) {
	slackUsers, err := getSlackUsersFromCache()
	if err != nil {
		return &models.SlackUser{}, err
	}
	for _, user := range slackUsers {
		if user.Name == name {
			return user, nil
		}
	}

	slackUsers, err = fetchSlackUsers()
	for _, user := range slackUsers {
		if user.Name == name {
			return user, nil
		}
	}
	return &models.SlackUser{}, errors.New("Slack id not found")
}

// ListSlackUsers returns slack user list
func ListSlackUsers() ([]*models.SlackUser, error) {
	cached, err := getSlackUsersFromCache()
	if len(cached) > 0 {
		return cached, err
	}
	return fetchSlackUsers()
}

func getSlackUsersFromCache() ([]*models.SlackUser, error) {
	file, err := os.Open(cacheSlackUserFilePath())
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

func putSlackUsersToCache(slackUsers []*models.SlackUser) error {
	file, err := os.OpenFile(cacheSlackUserFilePath(), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
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

func fetchSlackUsers() ([]*models.SlackUser, error) {
	token := os.Getenv("SLACK_API_TOKEN")
	if token == "" {
		log.Fatalf("You need to pass SLACK_API_TOKEN as environment variable.")
	}
	requestURL := SlackUserListURL + "?token=" + token
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

	putErr := putSlackUsersToCache(res)
	if putErr != nil {
		log.Println(putErr)
	}
	return res, nil
}
