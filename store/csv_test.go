package store

import (
	"crypto/rand"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/joho/godotenv"
	"github.com/wantedly/slack-mention-converter/models"
)

var (
	tmpdir string
)

func loadEnv() error {
	err := godotenv.Load(".env")
	if err != nil {
		return err
	}

	return nil
}

func sameUser(u1, u2 *models.User) bool {
	return u1.LoginName == u2.LoginName && u1.SlackName == u2.SlackName
}

func randomString() string {
	b := make([]byte, 8)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func TestUserFlow(t *testing.T) {
	csv := NewCSV(tmpdir)

	users, err := csv.ListUsers()
	if err != nil {
		t.Errorf("%v", err)
	}
	if len(users) != 0 {
		t.Fatalf("0 users should be present in initial state")
	}

	err = csv.AddUser(&models.User{LoginName: "awakia", SlackName: "naoyoshi"})
	if err != nil {
		t.Errorf("%v", err)
	}

	err = csv.AddUser(&models.User{LoginName: "kawasy", SlackName: "yoshi"})
	if err != nil {
		t.Errorf("%v", err)
	}

	err = csv.AddUser(&models.User{LoginName: "awakia", SlackName: "nao"})
	if err != nil {
		t.Errorf("%v", err)
	}

	users, err = csv.ListUsers()
	if err != nil {
		t.Errorf("%v", err)
	}
	if len(users) != 2 {
		t.Fatalf("2 users should be present after adding users but %v users", len(users))
	}

	awakia := &models.User{LoginName: "awakia", SlackName: "nao"}
	kawasy := &models.User{LoginName: "kawasy", SlackName: "yoshi"}

	if !sameUser(users[0], awakia) {
		t.Errorf("users[0] should be :%#v, but: %#v", awakia, users[0])
	}
	if !sameUser(users[1], kawasy) {
		t.Errorf("users[1] should be :%#v, but: %#v", kawasy, users[1])
	}

	user, err := csv.GetUser("awakia")
	if err != nil {
		t.Errorf("%v", err)
	}

	if !sameUser(user, awakia) {
		t.Errorf("user should be :%#v, but: %#v", awakia, user)
	}
}

// pending
func TestListSlackUser(t *testing.T) {
	t.Skip("TestListSlackUser is pending")

	csv := NewCSV(tmpdir)

	if err := loadEnv(); err != nil {
		t.Fatalf(err.Error())
	}

	slackUsers, err := csv.ListSlackUsers()
	if err != nil {
		t.Errorf("%v", err)
	}
	if len(slackUsers) < 10 {
		t.Fatalf("Data is less than 10, expected a much more larger number")
	}
	t.Logf("%v", slackUsers[0].ID)
	t.Logf("%v", slackUsers[0].Name)
}

func TestMain(m *testing.M) {
	tmpdir = filepath.Join(os.TempDir(), randomString())
	os.MkdirAll(tmpdir, 0755)
	defer os.RemoveAll(tmpdir)

	os.Exit(m.Run())
}
