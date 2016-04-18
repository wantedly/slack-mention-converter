package service

import (
	"testing"
)

func TestListSlackUser(t *testing.T) {
	slackUsers, err := ListSlackUsers()
	if err != nil {
		t.Errorf(err)
	}
	if len(slackUsers) < 10 {
		t.Fatalf("Data is less than 10, expected a much more larger number")
	}
	t.Logf("%v", slackUsers[0].ID)
	t.Logf("%v", slackUsers[0].Name)
}
