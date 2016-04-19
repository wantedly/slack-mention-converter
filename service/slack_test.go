package service

import (
	"log"
	"testing"

	"github.com/joho/godotenv"
)

func TestListSlackUser(t *testing.T) {
	loadEnv()
	slackUsers, err := ListSlackUsers()
	if err != nil {
		t.Errorf("%v", err)
	}
	if len(slackUsers) < 10 {
		t.Fatalf("Data is less than 10, expected a much more larger number")
	}
	t.Logf("%v", slackUsers[0].ID)
	t.Logf("%v", slackUsers[0].Name)
}

func loadEnv() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
